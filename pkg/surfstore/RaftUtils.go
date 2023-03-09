package surfstore

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"

	"google.golang.org/grpc"
)

type RaftConfig struct {
	RaftAddrs  []string
	BlockAddrs []string
}

func LoadRaftConfigFile(filename string) (cfg RaftConfig) {
	configFD, e := os.Open(filename)
	if e != nil {
		log.Fatal("Error Open config file:", e)
	}
	defer configFD.Close()

	configReader := bufio.NewReader(configFD)
	decoder := json.NewDecoder(configReader)

	if err := decoder.Decode(&cfg); err == io.EOF {
		return
	} else if err != nil {
		log.Fatal(err)
	}
	return
}

func NewRaftServer(id int64, config RaftConfig) (*RaftSurfstore, error) {
	// TODO Any initialization you need here

	isLeaderMutex := sync.RWMutex{}
	isCrashedMutex := sync.RWMutex{}

	peersInfo := []*PeerInfo{}
	var selfAddr string
	for i := 0; i < len(config.RaftAddrs); i++ {
		if int64(i) == id {
			selfAddr = config.RaftAddrs[i]
			continue
		}
		newInfo := PeerInfo{
			serverId:   int64(i),
			addr:       config.RaftAddrs[i],
			nextIndex:  0,
			matchIndex: -1,
		}
		peersInfo = append(peersInfo, &newInfo)
	}

	server := RaftSurfstore{
		isLeader:       false,
		isLeaderMutex:  &isLeaderMutex,
		term:           0,
		metaStore:      NewMetaStore(config.BlockAddrs),
		log:            make([]*UpdateOperation, 0),
		isCrashed:      false,
		isCrashedMutex: &isCrashedMutex,

		// newly added params
		peerInfo:    peersInfo,
		serverId:    id,
		commitIndex: -1,
		addr:        selfAddr,
	}

	return &server, nil
}

// TODO Start up the Raft server and any services here
func ServeRaftServer(server *RaftSurfstore) error {
	grpcServer := grpc.NewServer()
	RegisterRaftSurfstoreServer(grpcServer, server)

	lis, err := net.Listen("tcp", server.addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}
