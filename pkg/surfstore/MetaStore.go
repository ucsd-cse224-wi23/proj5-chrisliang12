package surfstore

import (
	context "context"
	"sync"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type MetaStore struct {
	FileMetaMap        map[string]*FileMetaData
	BlockStoreAddrs    []string
	ConsistentHashRing *ConsistentHashRing
	mutex              sync.Mutex
	UnimplementedMetaStoreServer
}

func (m *MetaStore) GetBlockStoreMap(ctx context.Context, blockHashesIn *BlockHashes) (*BlockStoreMap, error) {
	blockStoreMap := make(map[string]*BlockHashes)
	for _, hash := range blockHashesIn.Hashes {
		respServer := m.ConsistentHashRing.GetResponsibleServer(hash)

		if blockHashes, ok := blockStoreMap[respServer]; ok {
			blockHashes.Hashes = append(blockHashes.Hashes, hash)
			continue
		}
		newHashList := []string{}
		newHashList = append(newHashList, hash)
		blockStoreMap[respServer] = &BlockHashes{Hashes: newHashList}

	}
	return &BlockStoreMap{BlockStoreMap: blockStoreMap}, nil
}

func (m *MetaStore) GetBlockStoreAddrs(ctx context.Context, _ *emptypb.Empty) (*BlockStoreAddrs, error) {
	return &BlockStoreAddrs{BlockStoreAddrs: m.BlockStoreAddrs}, nil
}

func (m *MetaStore) GetFileInfoMap(ctx context.Context, _ *emptypb.Empty) (*FileInfoMap, error) {
	return &FileInfoMap{FileInfoMap: m.FileMetaMap}, nil
}

func (m *MetaStore) UpdateFile(ctx context.Context, fileMetaData *FileMetaData) (*Version, error) {
	fileName := fileMetaData.Filename
	version := fileMetaData.Version
	m.mutex.Lock()
	d, ok := m.FileMetaMap[fileName]
	if ok {
		if version == d.Version+1 {
			m.FileMetaMap[fileName] = fileMetaData
		} else {
			version = -1
		}
	} else {
		m.FileMetaMap[fileName] = fileMetaData
	}
	m.mutex.Unlock()
	return &Version{Version: version}, nil
}

// This line guarantees all method for MetaStore are implemented
var _ MetaStoreInterface = new(MetaStore)

func NewMetaStore(blockStoreAddrs []string) *MetaStore {
	return &MetaStore{
		FileMetaMap:        map[string]*FileMetaData{},
		BlockStoreAddrs:    blockStoreAddrs,
		ConsistentHashRing: NewConsistentHashRing(blockStoreAddrs),
	}
}
