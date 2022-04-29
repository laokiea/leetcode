package main

import "fmt"

// 剑指 Offer 03. 数组中重复的数字

func main() {
	fmt.Println(findRepeatNumber([]int{1, 3, 2, 4 ,2}))
}

func findRepeatNumber1(nums []int) int {
	for i := 0;i < len(nums);i++ {
		times := 0
		for _, n := range nums {
			if n == i {
				times++
				if times > 1 {
					return i
				}
			}
		}
	}
	return -1
}

// 第一次遇到数字x， 将其交换到数字x出
func findRepeatNumber2(nums []int) int {
	m := make(map[int]bool)
	for _, n := range nums {
		if ok, _ := m[n]; !ok {
			m[n] = true
		} else {
			return n
		}
	}
	return -1
}

// 第一次遇到数字x， 将其交换到数字x出
func findRepeatNumber(nums []int) int {
	for i, n := range nums {
		if i == n {
			continue
		} else {
			if n == nums[n] {
				return n
			} else {
				temp := nums[n]
				nums[n] = n
				nums[i] = temp
			}
		}
	}
	return -1
}