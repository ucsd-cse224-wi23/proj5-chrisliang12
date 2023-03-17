package surfstore

import (
	context "context"
	"log"
	"math"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type PeerInfo struct {
	infoMutex  *sync.Mutex
	serverId   int64
	addr       string
	nextIndex  int64
	matchIndex int64
	isFirstMsg bool
}

// TODO Add fields you need here
type RaftSurfstore struct {
	isLeader      bool
	isLeaderMutex *sync.RWMutex
	term          int64
	log           []*UpdateOperation

	metaStore *MetaStore

	peerInfo    []*PeerInfo
	serverId    int64
	commitIndex int64
	addr        string

	/*--------------- Chaos Monkey --------------*/
	isCrashed      bool
	isCrashedMutex *sync.RWMutex
	UnimplementedRaftSurfstoreServer
}

func (s *RaftSurfstore) GetFileInfoMap(ctx context.Context, empty *emptypb.Empty) (*FileInfoMap, error) {
	for {
		s.isCrashedMutex.RLock()
		isCrashed := s.isCrashed
		s.isCrashedMutex.RUnlock()
		if isCrashed {
			return nil, ERR_SERVER_CRASHED
		}

		s.isLeaderMutex.RLock()
		isLeader := s.isLeader
		s.isLeaderMutex.RUnlock()
		if !isLeader {
			return nil, ERR_NOT_LEADER
		}

		succ, _ := s.SendHeartbeat(ctx, empty)
		if succ.Flag {
			break
		}
	}

	return s.metaStore.GetFileInfoMap(ctx, empty)
}

func (s *RaftSurfstore) GetBlockStoreMap(ctx context.Context, hashes *BlockHashes) (*BlockStoreMap, error) {
	s.isCrashedMutex.RLock()
	isCrashed := s.isCrashed
	s.isCrashedMutex.RUnlock()
	if isCrashed {
		return nil, ERR_SERVER_CRASHED
	}

	s.isLeaderMutex.RLock()
	isLeader := s.isLeader
	s.isLeaderMutex.RUnlock()
	if !isLeader {
		return nil, ERR_NOT_LEADER
	}

	return s.metaStore.GetBlockStoreMap(ctx, hashes)
}

func (s *RaftSurfstore) GetBlockStoreAddrs(ctx context.Context, empty *emptypb.Empty) (*BlockStoreAddrs, error) {
	s.isCrashedMutex.RLock()
	isCrashed := s.isCrashed
	s.isCrashedMutex.RUnlock()
	if isCrashed {
		return nil, ERR_SERVER_CRASHED
	}

	s.isLeaderMutex.RLock()
	isLeader := s.isLeader
	s.isLeaderMutex.RUnlock()
	if !isLeader {
		return nil, ERR_NOT_LEADER
	}

	return s.GetBlockStoreAddrs(ctx, empty)
}

func (s *RaftSurfstore) UpdateFile(ctx context.Context, filemeta *FileMetaData) (*Version, error) {
	log.Println(" ")
	log.Println("----------", s.serverId, " UpdateFile----------")

	// sanity check
	s.isCrashedMutex.RLock()
	isCrashed := s.isCrashed
	s.isCrashedMutex.RUnlock()
	if isCrashed {
		log.Println("--", s.serverId, " crashed")
		return nil, ERR_SERVER_CRASHED
	}

	s.isLeaderMutex.RLock()
	isLeader := s.isLeader
	s.isLeaderMutex.RUnlock()
	if !isLeader {
		log.Println("--", s.serverId, " not leader")
		return nil, ERR_NOT_LEADER
	}

	// NOTE: Check Version
	s.metaStore.mutex.Lock()
	if d, ok := s.metaStore.FileMetaMap[filemeta.Filename]; ok {
		if d.Version+1 != filemeta.Version {
			return &Version{Version: -1}, nil
		}
	}
	s.metaStore.mutex.Unlock()

	// create a new log entry; add it to the log
	newEntry := UpdateOperation{
		Term:         s.term,
		FileMetaData: filemeta,
	}

	s.log = append(s.log, &newEntry)
	log.Println(s.serverId, "Append new log entry, now the log are ", s.log)

	// send heartbeat to replicate log
	// if majority node fail, block the operation until majority recover
	// NOTE: This round of SendingHeartbeat is for log replication (not commitment)
	for {
		s.isLeaderMutex.RLock()
		isLeader := s.isLeader
		s.isLeaderMutex.RUnlock()
		if !isLeader {
			return &Version{Version: -1}, ERR_NOT_LEADER
		}

		s.isCrashedMutex.RLock()
		isCrashed := s.isCrashed
		s.isCrashedMutex.RUnlock()
		if isCrashed {
			return &Version{Version: -1}, ERR_SERVER_CRASHED
		}

		succ, err := s.SendHeartbeat(ctx, &emptypb.Empty{})
		if err != nil || !succ.Flag {
			return &Version{Version: -1}, err
		}

		if succ.Flag && len(s.log)-1 > int(s.commitIndex) {
			for i := s.commitIndex + 1; i < int64(len(s.log)); i++ {
				if s.log[i].Term == s.term {
					s.commitIndex = i
				}
			}
			break
		}
	}

	// NOTE: Commit since majority of nodes have the log
	v, _ := s.metaStore.UpdateFile(ctx, filemeta)

	// NOTE: Send another round of heartbeat to commit the log
	log.Println(" ")
	log.Println("----------", s.serverId, " UpdateFile: commit msg start sending----------")

	_, err := s.SendHeartbeat(ctx, &emptypb.Empty{})
	if err != nil {
		log.Println("serverId: ", s.serverId, "commit fail")
		return v, err
	}

	return v, nil
}

// 1. Reply false if term < currentTerm (§5.1)
// 2. Reply false if log doesn’t contain an entry at prevLogIndex whose term
// matches prevLogTerm (§5.3)
// 3. If an existing entry conflicts with a new one (same index but different
// terms), delete the existing entry and all that follow it (§5.3)
// 4. Append any new entries not already in the log
// 5. If leaderCommit > commitIndex, set commitIndex = min(leaderCommit, index
// of last new entry)
func (s *RaftSurfstore) AppendEntries(ctx context.Context, input *AppendEntryInput) (*AppendEntryOutput, error) {
	// WARN: comment out the log before submission
	res := AppendEntryOutput{
		Success:      false,
		MatchedIndex: -1,
		Term:         s.term,
		ServerId:     s.serverId,
	}

	s.isCrashedMutex.RLock()
	isCrashed := s.isCrashed
	s.isCrashedMutex.RUnlock()
	if isCrashed {
		return &res, ERR_SERVER_CRASHED
	}

	// 1. reply false if term < currTerm
	if input.Term < s.term {
		return &res, nil
	}

	if input.Term > s.term {
		s.term = input.Term
		res.Term = s.term
		s.isLeaderMutex.RLock()
		isLeader := s.isLeader
		s.isLeaderMutex.RUnlock()
		if isLeader {
			s.isLeaderMutex.Lock()
			s.isLeader = false
			s.isLeaderMutex.Unlock()
		}
	}

	// 2. Reply false if log doesn’t contain an entry at prevLogIndex whose term
	// matches prevLogTerm (§5.3)
	if input.PrevLogIndex != -1 && int64(len(s.log)-1) < input.PrevLogIndex {
		return &res, nil
	} else if input.PrevLogIndex != -1 && int64(len(s.log)-1) > input.PrevLogIndex {
		if s.log[input.PrevLogIndex].Term != input.PrevLogTerm {
			return &res, nil
		}
	}

	// 3. If an existing entry conflicts with a new one (same index but different
	// terms), delete the existing entry and all that follow it (§5.3)
	// if input.PrevLogIndex != -1 || len(input.Entries) > 0 {
	// 	s.log = s.log[:input.PrevLogIndex+1]
	// }
	s.log = s.log[:input.PrevLogIndex+1]

	// 4. Append any new entries not already in the log
	if len(input.Entries) > 0 {
		s.log = append(s.log, input.Entries...)
	}

	// 5. If leaderCommit > commitIndex, set commitIndex = min(leaderCommit, index
	// of last new entry)
	if input.LeaderCommit > s.commitIndex {
		currCommitIdx := s.commitIndex
		s.commitIndex = int64(math.Min(float64(input.LeaderCommit), float64(len(s.log)-1)))

		for i := currCommitIdx + 1; i <= s.commitIndex; i++ {
			s.metaStore.UpdateFile(ctx, s.log[i].FileMetaData)
		}
	}

	res.Success = true
	res.MatchedIndex = int64(len(s.log) - 1)

	return &res, nil
}

func (s *RaftSurfstore) SetLeader(ctx context.Context, _ *emptypb.Empty) (*Success, error) {
	// WARN: comment out the log before submission
	log.Println("")
	log.Println("----------Set Leader to ", s.serverId, "----------")
	// sanity check
	s.isCrashedMutex.RLock()
	isCrashed := s.isCrashed
	s.isCrashedMutex.RUnlock()
	if isCrashed {
		log.Println(s.serverId, " crashed when setting to leader")
		return &Success{Flag: false}, ERR_SERVER_CRASHED
	}

	// set leader
	s.isLeaderMutex.Lock()
	s.isLeader = true
	s.isLeaderMutex.Unlock()
	s.term++

	// setup initial PeerInfo
	logCount := len(s.log)
	for _, info := range s.peerInfo {
		info.infoMutex.Lock()
		info.matchIndex = -1
		info.nextIndex = int64(logCount)
		info.isFirstMsg = true
		info.infoMutex.Unlock()
	}

	log.Println(s.serverId, "curr log: ", s.log)
	log.Println("----------Set Leader End ", s.serverId, "----------")

	return &Success{Flag: true}, nil
}

func (s *RaftSurfstore) SendHeartbeat(ctx context.Context, _ *emptypb.Empty) (*Success, error) {
	// sanity check
	s.isCrashedMutex.RLock()
	isCrashed := s.isCrashed
	s.isCrashedMutex.RUnlock()
	if isCrashed {
		return &Success{Flag: false}, ERR_SERVER_CRASHED
	}
	s.isLeaderMutex.RLock()
	isLeader := s.isLeader
	s.isLeaderMutex.RUnlock()
	if !isLeader {
		return &Success{Flag: false}, ERR_NOT_LEADER
	}

	// channel to receive AppendEntryOutput from goroutines
	nodeCount := len(s.peerInfo)
	res := make(chan *AppendEntryOutput, nodeCount)
	var wg sync.WaitGroup
	wg.Add(nodeCount)
	// send heartbeat in parallel
	for _, info := range s.peerInfo {
		go s.handleSendingHeartbeat(info, res, &wg)
	}

	// check the response
	isMajAlive := false
	aliveServerNum := 1

	wg.Wait()

	for i := 0; i < nodeCount; i++ {
		r := <-res
		if r != nil && r.Success {
			aliveServerNum++
		}
	}

	if float64(aliveServerNum) >= math.Ceil(float64(nodeCount+1)/2) {
		log.Println(s.serverId, " majority alive")
		isMajAlive = true
	}

	return &Success{Flag: isMajAlive}, nil
}

func (s *RaftSurfstore) handleSendingHeartbeat(peerInfo *PeerInfo, res chan *AppendEntryOutput, wg *sync.WaitGroup) {
	PrintPeerInfo(peerInfo)
	defer wg.Done()
	var rcvOutput *AppendEntryOutput

	conn, err := grpc.Dial(peerInfo.addr, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		res <- &AppendEntryOutput{ServerId: peerInfo.serverId, Term: 0, Success: false, MatchedIndex: -1}
		return
	}
	client := NewRaftSurfstoreClient(conn)

	for {
		// set prevLogTerm and Index based on peerInfo
		prevLogIndex := peerInfo.nextIndex - 1
		prevLogTerm := int64(0)
		if prevLogIndex != -1 {
			prevLogTerm = s.log[prevLogIndex].Term
		}

		// set entry if there are any
		var entry []*UpdateOperation
		if !peerInfo.isFirstMsg && len(s.log)-1 != int(peerInfo.matchIndex) {
			entry = s.log[peerInfo.matchIndex+1:]
		}
		currEntry := AppendEntryInput{
			Term:         s.term,
			PrevLogIndex: prevLogIndex,
			PrevLogTerm:  prevLogTerm,
			Entries:      entry,
			LeaderCommit: s.commitIndex,
		}
		PrintAppendEntryInput(&currEntry, s.serverId, peerInfo.serverId)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		op, err := client.AppendEntries(ctx, &currEntry)
		peerInfo.isFirstMsg = false
		if err != nil {
			if strings.Contains(err.Error(), ERR_SERVER_CRASHED.Error()) {
				rcvOutput = op
				break
			}
		}

		if err == nil {
			PrintAPpendEntryOutput(op, s.serverId, peerInfo.serverId)

			// if the resp has a higher term, which means there is a new isLeader
			// set the isLeader to false and return
			if op.Term > s.term {
				s.isLeaderMutex.Lock()
				s.isLeader = false
				s.isLeaderMutex.Unlock()
				rcvOutput = op
				break
			}

			// update peerInfo
			if op.MatchedIndex > peerInfo.matchIndex {
				peerInfo.infoMutex.Lock()
				peerInfo.matchIndex = op.MatchedIndex
				peerInfo.nextIndex = op.MatchedIndex + 1
				peerInfo.infoMutex.Unlock()
			}

			// if success, break the loop and return
			if op.Success {
				rcvOutput = op
				break
			}

			// if there is no MatchedIndex, decline the prevLogIndex
			if op.MatchedIndex == -1 {
				peerInfo.infoMutex.Lock()
				peerInfo.nextIndex--
				peerInfo.matchIndex = -1
				peerInfo.infoMutex.Unlock()
			}
		}
	}
	res <- rcvOutput
}

// ========== DO NOT MODIFY BELOW THIS LINE =====================================

func (s *RaftSurfstore) Crash(ctx context.Context, _ *emptypb.Empty) (*Success, error) {
	s.isCrashedMutex.Lock()
	s.isCrashed = true
	s.isCrashedMutex.Unlock()

	return &Success{Flag: true}, nil
}

func (s *RaftSurfstore) Restore(ctx context.Context, _ *emptypb.Empty) (*Success, error) {
	s.isCrashedMutex.Lock()
	s.isCrashed = false
	s.isCrashedMutex.Unlock()

	return &Success{Flag: true}, nil
}

func (s *RaftSurfstore) GetInternalState(ctx context.Context, empty *emptypb.Empty) (*RaftInternalState, error) {
	fileInfoMap, _ := s.metaStore.GetFileInfoMap(ctx, empty)
	s.isLeaderMutex.RLock()
	state := &RaftInternalState{
		IsLeader: s.isLeader,
		Term:     s.term,
		Log:      s.log,
		MetaMap:  fileInfoMap,
	}
	s.isLeaderMutex.RUnlock()

	return state, nil
}

var _ RaftSurfstoreInterface = new(RaftSurfstore)
