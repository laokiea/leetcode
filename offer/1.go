package main

import "fmt"

//剑指 Offer II 001. 整数除法
// 思路1：减法代替除法
// 1. 15/2 => 不断的15 - 2，判断结果是否大于2， 第7次时，3-2=1，1<2 结束，所以结果为1
// 2. 32位有整形，最大整为2^31-1，最小值为-2^31,所以负数不会溢出，a,b只要大于0，转成负数
// 3. -2^31 / -1 = 2^31，溢出

func main() {
	fmt.Println(solution2(15,2))
}

func solution1(a,b int) int {
	IntMax := int32(^uint32(0) >> 1)
	IntMin := ^IntMax
	var result int
	var flag bool
	if int32(a) == IntMin && b == -1 {
		return 0x7FFFFFFF
	}
	if a >= 0 && b <= 0 {
		flag = true
		a = -a
	}
	if a <= 0 && b >= 0 {
		flag = true
		b = -b
	}
	if a >= 0 && b >= 0 {
		a = -a
		b = -b
	}
	if a > b {
		return 0
	}
	for a <= b {
		a -= b
		result++
	}
	if flag {
		return -result
	} else {
		return result
	}
}

func solution2(a, b int) int {
	IntMax := int32(^uint32(0) >> 1)
	IntMin := ^IntMax
	var result int
	var flag bool
	if int32(a) == IntMin && b == -1 {
		return 0x7FFFFFFF
	}
	if a >= 0 && b <= 0 {
		flag = true
		a = -a
	}
	if a <= 0 && b >= 0 {
		flag = true
		b = -b
	}
	if a >= 0 && b >= 0 {
		a = -a
		b = -b
	}
	if a > b {
		return 0
	}
	//-15 / -2
	for b >= a {
		loop := 1
		c := b
		for c + c > a {
			loop <<= 1
			c <<= 1
		}
		a -= c
		result += loop
	}
	if flag {
		return -result
	} else {
		return result
	}
}