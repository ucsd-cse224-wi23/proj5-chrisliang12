package surfstore

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
)

type Node struct {
	next *Node
	hash string
}

type ConsistentHashRing struct {
	ServerMap map[string]string
	RingHead  *Node
}

func (c ConsistentHashRing) GetResponsibleServer(blockId string) string {
	ptr := c.RingHead
	for i := 0; i < len(c.ServerMap); i++ {
		if blockId < ptr.hash {
			break
		}
		ptr = ptr.next
	}
	return c.ServerMap[ptr.hash]
}

func (c ConsistentHashRing) Hash(addr string) string {
	h := sha256.New()
	h.Write([]byte(addr))
	return hex.EncodeToString(h.Sum(nil))
}

func NewConsistentHashRing(serverAddrs []string) *ConsistentHashRing {
	var cstHashRing ConsistentHashRing
	serverMap := make(map[string]string)
	hashes := []string{}

	for _, serverAddr := range serverAddrs {
		currServerName := "blockstore" + serverAddr
		serverHash := cstHashRing.Hash(currServerName)
		serverMap[serverHash] = serverAddr
		hashes = append(hashes, serverHash)
	}
	sort.Strings(hashes)

	head := Node{hash: hashes[0]}
	nodePtr := &head
	for _, hash := range hashes[1:] {
		newNode := Node{hash: hash}
		nodePtr.next = &newNode
		nodePtr = nodePtr.next
	}
	nodePtr.next = &head

	cstHashRing = ConsistentHashRing{ServerMap: serverMap, RingHead: &head}
	return &cstHashRing
}
