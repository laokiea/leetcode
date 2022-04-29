package main

import "fmt"

//剑指 Offer II 006. 排序数组中两个数字之和
// 数组已经排序
// [2, 4, 6, 8, 10] 12 [0,4]

// 思路：头尾双指针

func main() {
	fmt.Println(twoSum([]int{0,0,3,4}, 0))
}

func twoSum(numbers []int, target int) []int {
	head := 0
	tail := len(numbers) - 1
	for numbers[head] <= numbers[tail] {
		if numbers[head] + numbers[tail] == target {
			return []int{head, tail}
		} else if numbers[head] + numbers[tail] > target {
			tail--
		} else {
			head++
		}
	}
	return []int{0}
}