package main

import "fmt"

//剑指 Offer II 003. 前 n 个数字二进制中 1 的个数
// 偶数n的二进制表示中，1的个数一定和n/2的二进制表示中1的个数相同，因为最后一位肯定为0，右移一位没影响

func main() {
	fmt.Println(countBits(5))
}

//00000000 00000000 00000000 00000110
//10000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000

func countBits1(n uint64) []uint64 {
	r := make([]uint64, n)
	var i uint64
	for ;i <= n;i++ {
		var t uint64
		for j := 0;j < 64;j++ {
			if ((i >> j) << 63) == 1<<63 {
				t++
			}
		}
		r = append(r, t)
	}
	return r
}

func countBits2(n int) []int {
	r := make([]int, 0)
	var i uint32
	for ;int(i) <= n;i++ {
		var t int
		var j uint32
		for j = 0;j < 64;j++ {
			if ((i >> j) << 31) == 1<<31 {
				t++
			}
		}
		r = append(r, t)
	}
	return r
}

func countBits(n int) []int {
	r := make([]int, n+1)
	for i := 0;i <= n;i++ {
		if i&1 == 0 {
			r[i] = r[i>>1]
		} else {
			r[i] = r[i-1]+1
		}
	}
	return r
}