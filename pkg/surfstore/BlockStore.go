package surfstore

import (
	context "context"
	"fmt"
	"sync"

	"google.golang.org/protobuf/types/known/emptypb"
)

type BlockStore struct {
	BlockMap map[string]*Block
	mutex    sync.Mutex
	UnimplementedBlockStoreServer
}

func (bs *BlockStore) GetBlockHashes(ctx context.Context, _ *emptypb.Empty) (*BlockHashes, error) {
	hashes := []string{}
	for _, b := range bs.BlockMap {
		hash := GetBlockHashString(b.BlockData)
		hashes = append(hashes, hash)
	}
	return &BlockHashes{Hashes: hashes}, nil
}

func (bs *BlockStore) GetBlock(ctx context.Context, blockHash *BlockHash) (*Block, error) {
	bs.mutex.Lock()
	b, ok := bs.BlockMap[blockHash.Hash]
	if !ok {
		bs.mutex.Unlock()
		return nil, fmt.Errorf("cannot get block")
	}
	bs.mutex.Unlock()
	return b, nil
}

func (bs *BlockStore) PutBlock(ctx context.Context, block *Block) (*Success, error) {
	hash := GetBlockHashString(block.BlockData)
	bs.mutex.Lock()
	bs.BlockMap[hash] = block
	bs.mutex.Unlock()
	return &Success{Flag: true}, nil
}

// Given a list of hashes “in”, returns a list containing the
// subset of in that are stored in the key-value store
func (bs *BlockStore) HasBlocks(
	ctx context.Context,
	blockHashesIn *BlockHashes,
) (*BlockHashes, error) {
	hashout := make([]string, 0)
	for _, hash := range blockHashesIn.Hashes {
		if _, ok := bs.BlockMap[hash]; !ok {
			continue
		}
		hashout = append(hashout, hash)
	}
	if len(hashout) == 0 {
		return nil, fmt.Errorf("cannot find blocks")
	}
	return &BlockHashes{Hashes: hashout}, nil
}

// This line guarantees all method for BlockStore are implemented
var _ BlockStoreInterface = new(BlockStore)

func NewBlockStore() *BlockStore {
	return &BlockStore{
		BlockMap: map[string]*Block{},
	}
}
