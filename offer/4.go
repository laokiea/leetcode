package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{0,1,0,1,0,1,100}))
	//fmt.Println(singleNumber([]int{3,3,5,3}))
}

// [2, 2, 3, 2]

func singleNumber1(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = 1
		} else {
			if m[n] == 2 {
				delete(m, n)
			} else {
				m[n] = m[n] + 1
			}
		}
	}
	for k := range m {
		return k
	}
	return -1
}

func singleNumber(nums []int) int {
	var r int32
	var i int32
	for ;i < 32;i++ {
		var s int32
		for _, n := range nums {
			t := int32(n)
			s += (t >> i) & 1
		}
		r += (s % 3) << i
	}
	return int(r)
}