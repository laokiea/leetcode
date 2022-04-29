package main

import (
	"fmt"
)

type Node struct {
	v interface{}
	nxt *Node
}

// 1 2 3 4 5 nil
// prev = nil
// head = 1
// cur = head = 1

// nxt = cur.next = 2
// cur.next = prev => 1.next = nil
// prev = cur = 1
// cur = nxt = 2
// 1 nil 2 3 4 5 nil

// cur = 2
// nxt = cur.next = 3
// cur.next = prev => 2.next = 1
// prev = cur
// cur = nxt
// 2 1 nil 3 4 5 nil

func main() {
	n := construct()
	dump(n)
	dump(reverse(n))
}

func construct() *Node {
	var cur *Node
	for i := 5;i > 0;i-- {
		node := Node{i, cur}
		cur = &node
	}
	return cur
}

func dump(node *Node) {
	for node != nil {
		fmt.Println(node.v.(int))
		node = node.nxt
	}
}

// 1 2 3 4 5
func reverse(node *Node) *Node {
	var cur, nxt, prev *Node
	cur = node
	for cur != nil {
		nxt = cur.nxt
		cur.nxt = prev
		prev = cur
		cur = nxt
	}
	return prev
}