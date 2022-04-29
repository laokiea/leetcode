package main

import (
	"fmt"
)

type Heap []int

// left: 2*i + 1
// right 2*i + 2

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Insert(v int) {
	*h = append(*h, v)
}

func (h Heap) Push(v int) {
	(*Heap).Insert(&h, v)
	if len(h) == 1 {
		return
	}
	h.up(len(h) - 1)
}

func (h *Heap) Remove(i int) int {
	l := len(*h)
	if i < 0 || i > l-1 {
		return -1
	}
	// 把待删除的节点和最后一个节点互换
	h.Swap(i, l-1)
	x := (*h)[l-1]
	*h = (*h)[:l-1]
	// 现在第i个节点是之前最后一个节点
	// 需要比较和父节点的大小关系
	// 如果比父节点大 就down， 比父节点小就up
	var p int
	if i%2 == 0 {
		p = (i - 2)/2
	} else {
		p = (i - 1)/2
	}
	if p <= 0 || h.Less(p, i) {
		h.down(i)
	} else {
		h.up(i)
	}
	return x
}

func (h *Heap) Pop() int {
	return h.Remove(0)
}

func Build(arr []int) Heap {
	h := Heap(arr)
	// 从第一个非叶子节点开始
	// 最后一个节点 求父节点 即第一个非叶子节点
	l := len(h) - 1
	var firstNoLeafNode int
	if (l - 1)%2 == 0 {
		firstNoLeafNode = (l-2)/2
	} else {
		firstNoLeafNode = (l-1)/2
	}
	for i:=firstNoLeafNode;i>=0;i-- {
		h.down(i)
	}
	return h
}

func (h Heap) up(i int) {
	for {
		p := (i - 1) / 2
		if h.Less(p, i) || i == p {
			break
		}
		h.Swap(i, p)
		i = p
	}
}

func (h Heap) down(i int) {
	for {
		// 和最小的子节点比较，交换，直到没有字节点为止
		left := 2*i + 1
		right := 2*i + 2
		// 新一轮的左节点大于最长长度，说明本身已经是子节点
		if left >= len(h) {
			break
		}
		minChild := left
		if right < len(h) && h.Less(right, left) {
			minChild = right
		}
		if h.Less(i, minChild) {
			break
		}
		h.Swap(minChild, i)
		i = minChild
	}
}

func (h Heap) sort() []int {
	sortArr := make([]int, 0)
	for len(h) > 0 {
		sortArr = append(sortArr, h.Pop())
		fmt.Println(sortArr)
	}
	return sortArr
}

func main() {
	arr := []int{4,8,7,6,5,3,1,11,9,17}
	h := Build(arr)
	fmt.Println(h.sort())
}