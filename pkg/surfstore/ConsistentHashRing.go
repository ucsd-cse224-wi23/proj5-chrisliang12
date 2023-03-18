package surfstore

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
)

type ConsistentHashRing struct {
	ServerMap map[string]string
	HashRing  []*string
}

// use binary search to find the Responsible Server
func (c ConsistentHashRing) GetResponsibleServer(blockId string) string {
	l, r := 0, len(c.HashRing)
	for l < r {
		mid := l + (r-l)/2
		if *(c.HashRing[mid]) < blockId {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return c.ServerMap[*(c.HashRing[l])]
}

func (c ConsistentHashRing) Hash(addr string) string {
	h := sha256.New()
	h.Write([]byte(addr))
	return hex.EncodeToString(h.Sum(nil))
}

type sortableSlice []*string

func (s sortableSlice) Len() int {
	return len(s)
}

func (s sortableSlice) Less(i, j int) bool {
	return *s[i] < *s[j]
}

func (s sortableSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func NewConsistentHashRing(serverAddrs []string) *ConsistentHashRing {
	var cstHashRing ConsistentHashRing
	serverMap := make(map[string]string)
	hashes := []*string{}

	for _, serverAddr := range serverAddrs {
		currServerName := "blockstore" + serverAddr
		serverHash := cstHashRing.Hash(currServerName)
		serverMap[serverHash] = serverAddr
		hashes = append(hashes, &serverHash)
	}
	sort.Sort(sortableSlice(hashes))

	cstHashRing = ConsistentHashRing{ServerMap: serverMap, HashRing: hashes}
	return &cstHashRing
}
