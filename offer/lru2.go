package main

import "fmt"

func main() {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Get(1)
	lruCache.Put(3, 3)
	lruCache.Get(2)
	lruCache.Put(4, 4)
	//lruCache.Get(1)
	//lruCache.Get(3)
	//lruCache.Get(4)
	lruCache.Debug()
}

type LRUCache struct {
	Capacity int
	Entries map[int]*HashEntry
	Used int
	TopNode *Node
	BtmNode *Node
}

type Node struct {
	Next *Node
	Prev *Node
	Value int
	Key int
}

type HashEntry struct {
	Key int
	Node *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache {
		Capacity: capacity,
		Entries: make(map[int]*HashEntry),
		TopNode: NewNode(-1, -1),
		BtmNode: NewNode(-1, -1),
	}
	cache.TopNode.Next = cache.BtmNode
	cache.BtmNode.Prev = cache.TopNode
	return cache
}

func NewNode(key int, value int) *Node {
	return &Node{nil,nil,value,key}
}

func (this *LRUCache) Debug() {
	node := this.BtmNode
	for node != nil {
		fmt.Printf("%p %v\n", node, node)
		node = node.Prev
	}
}

func (this *LRUCache) Get(key int) int {
	if this.Entries[key] == nil {
		return -1
	} else {
		v := this.Entries[key].Node.Value
		this.Put(key, v)
		return v
	}
}


func (this *LRUCache) Put(key int, value int)  {
	//
	entry := &HashEntry{
		Node: NewNode(key, value),
		Key: key,
	}
	if this.Entries[key] == nil {
		this.Entries[key] = entry
		// 新元素移动到top
		this.addToTop(entry.Node)
		this.Used++
		if this.Used > this.Capacity {
			// 删除最后一个
			this.removeTail()
		}
	} else {
		node := this.Entries[key].Node
		node.Value = value
		this.moveToTop(node)
	}
}

// A = B = C
// A.Prev = B.Prev
// C.Next = B.Next
func (this *LRUCache) removeTail() {
	delete(this.Entries, this.BtmNode.Prev.Key)
	this.BtmNode.Prev.Prev.Next = this.BtmNode
	this.BtmNode.Prev = this.BtmNode.Prev.Prev
	this.Used--
}

func (this *LRUCache) moveToTop(node *Node) {
	// [(0,0), (0,0)]
	// [(0,0), (0,0), (1,2)]
	this.remove(node)
	this.addToTop(node)
}

func (this *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// A = B (c)
// C.Next = A
// c.Prev = A.Prev
// A.Prev.Next = C
// A.Prev = C
func (this *LRUCache) addToTop(node *Node) {
	node.Prev = this.TopNode
	node.Next = this.TopNode.Next
	this.TopNode.Next.Prev = node
	this.TopNode.Next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */