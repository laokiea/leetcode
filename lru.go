package main

import (
	"fmt"
)

// []
// capacity = 4
// insert (1, 2) : [(1, 2)]
// insert (2, 4) : [(1, 2), (2, 4)]
// insert (3, 8) : [(1, 2), (2, 4), (3, 8)]
// insert (4, 16) : [(1, 2), (2, 4), (3, 8), (4, 16)]
// get (2) : [(1, 2), (3, 8), (4, 16), (2, 4)]
// insert (5, 32) : [(3, 8), (4, 16), (2, 4), (5, 32)]
// insert (2, 64) : [(3, 8), (4, 16), (5, 32), (2, 64)]
// insert (6, 128) : [(4, 16), (5, 32), (2, 64), (6,128)]
func main() {
	lruCache := Constructor(1)
	lruCache.Put(2, 1) // [(2, 1)]
	lruCache.Get(2) // [(2, 1)]
	lruCache.Put(3, 2) // [(3, 2)]
	lruCache.Get(2)  // [(3, 2)]
	lruCache.Get(3) // [(3, 2)]
	lruCache.Debug()
	//fmt.Println(lruCache.BtmIdx, lruCache.TopIdx)
}

// LRU淘汰算法：Latest Recently Used
// 最近未使用的元素处于队尾

type LRUCache struct {
	Capacity int
	Entries map[uint64]*HashEntry
	HashFunc func(key int) uint64
	Node *ListNode
	Used int
	TopIdx uint64
	BtmIdx uint64
}

type ListNode struct {
	Next *ListNode
	Prev *ListNode
	Value int
	Key uint64
}

type HashEntry struct {
	Key int
	Node *ListNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Capacity: capacity,
		Entries: make(map[uint64]*HashEntry),
		HashFunc: HashFunc,
		Node: nil,
	}
	return cache
}

func (this *LRUCache) Debug() {
	node := this.Entries[this.BtmIdx].Node
	for node != nil {
		fmt.Printf("%p %v\n", node, node)
		node = node.Prev
	}
}

// []
// insert (1, 2) : [(1, 2)]
// insert (2, 4) : [(1, 2), (2, 4)]
// insert (3, 8) : [(1, 2), (2, 4), (3, 8)]
// insert (4, 16) : [(1, 2), (2, 4), (3, 8), (4, 16)]
// get (2) : [(1, 2), (3, 8), (4, 16), (2, 4)]

func (this *LRUCache) Get(key int) int {
	idx := this.GetHashIdx(key)
	if this.Entries[idx] == nil {
		return -1
	} else {
		elemNode := this.Entries[idx].Node
		this.moveElementToTop(elemNode)
		return elemNode.Value
	}
}

// []
// capacity = 4
// insert (1, 2) : [(1, 2)]
// insert (2, 4) : [(1, 2), (2, 4)]
// insert (3, 8) : [(1, 2), (2, 4), (3, 8)]
// insert (4, 16) : [(1, 2), (2, 4), (3, 8), (4, 16)]
// get (2) : [(1, 2), (3, 8), (4, 16), (2, 4)]
// insert (5, 32) : [(3, 8), (4, 16), (2, 4), (5, 32)]
// insert (2, 64) : [(3, 8), (4, 16), (5, 32), (2, 64)]
// insert (6, 128) : [(4, 16), (5, 32), (2, 64), (6,128)]

func (this *LRUCache) Put(key int, value int)  {
	idx := this.GetHashIdx(key)
	// key exists
	if this.Entries[idx] != nil {
		elemNode := this.Entries[idx].Node
		elemNode.Value = value
		this.moveElementToTop(elemNode)
	} else {
		node := &ListNode{
			nil,
			nil,
			value,
			uint64(key),
		}
		entry := &HashEntry{
			Key: key,
			Node: node,
		}
		if this.Used == 0 {
			this.BtmIdx = idx
			this.TopIdx = idx
			this.Entries[idx] = entry
			this.Used++
		} else {
			btmPrevNode := this.Entries[this.BtmIdx].Node.Prev
			entry.Node.Next = this.Entries[this.TopIdx].Node
			this.Entries[this.TopIdx].Node.Prev = node
			this.Entries[idx] = entry
			this.TopIdx = idx
			if this.Used == this.Capacity {
				//full, delete the oldest element
				this.Entries[this.BtmIdx] = nil
				if btmPrevNode != nil {
					btmPrevNode.Next = nil
					this.BtmIdx = btmPrevNode.Key
				} else {
					this.BtmIdx = idx
				}
			} else {
				this.Used++
			}
		}
	}
}

func (this *LRUCache) GetHashIdx(key int) uint64 {
	//hash := this.HashFunc(key)
	//idx := hash & uint64(this.Capacity)
	idx := uint64(key)
	return idx
}

func (this *LRUCache) moveElementToTop(elemNode *ListNode) {
	//  move element to the top
	if elemNode.Prev != nil {
		// let's move
		var endNode,elemNodeNext,elemNodePrev *ListNode
		endNode = elemNode
		for endNode.Prev != nil {
			endNode = endNode.Prev
		}
		// endNode is the first node
		endNode.Prev = elemNode
		elemNodeNext = elemNode.Next
		elemNodePrev = elemNode.Prev
		elemNode.Next = endNode
		elemNodePrev.Next = elemNodeNext
		if elemNodeNext != nil {
			elemNodeNext.Prev = elemNodePrev
		}
		elemNode.Prev = nil
		if this.BtmIdx == elemNode.Key {
			this.BtmIdx = endNode.Key
		}
		this.TopIdx = elemNode.Key
	}
}

func HashFunc(key int) uint64 {
	var hash uint64 = 5381
	keyStr := fmt.Sprint(key)
	keyStrLen := len(keyStr)
	for keyStrLen > 0 {
		hash = (hash << 5 + hash) + uint64(keyStr[len(keyStr) - keyStrLen])
	}
	return hash
}