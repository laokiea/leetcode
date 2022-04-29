package main

import "fmt"

//剑指 Offer II 009. 乘积小于 K 的子数组
//给定一个正整数数组 nums和整数 k ，请找出该数组内乘积小于 k 的连续的子数组的个数。

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10,5,2,6}, 100))
}

// [10,5,2,6]
func numSubarrayProductLessThanK(nums []int, k int) int {
	var left, right, result, total int
	total = 1
	for ;left < len(nums);left++ {
		right = left
		total *= nums[right]
		for total < k {
			fmt.Println(total)
			result++
			right++
			if right == len(nums) {
				break
			}
			total *= nums[right]
		}
		total = 1
	}
	return result
}

func min9(x, y int) int {
	if x > y {
		return y
	}
	return x
}