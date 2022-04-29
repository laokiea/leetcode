package main

import (
	"math"
)

// 剑指 Offer II 008. 和大于等于 target 的最短子数组
//给定一个含有 n 个正整数的数组和一个正整数 target 。
//
//找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
//

//输入：target = 7, nums = [2,3,1,2,4,3]

func main() {
	//fmt.Println(minSubArrayLen(3, []int{2,3,1,2,4,3}))
}

func minSubArrayLen(target int, nums []int) int {
	//[2,3,1,2,4,3]
	var left,right,total int
	var result = math.MaxInt32
	for ;right < len(nums);right++ {
		total += nums[right]
		for total >= target {
			result = min(result, right - left + 1)
			total -= nums[left]
			left++
		}
	}
	if result == math.MaxInt32 {
		result = 0
	}
	return result
}

//{2,3,1,2,4,3}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 3
// 2,3,1,2,4,3

//{2, 3} left: 0 right: 1 total: 5 result: 2
//{3} left: 1 right: 1 total: 5 result: 1
//{} left: 2 right: 1 total: 0 result: 1
//{1} left: 2 right: 2 total: 1 result: 1
//{1, 2} left: 2 right: 3 total: 3 result: 1
//{2} left: 3 right: 3 total: 3 result: 1
//{2, 4} left: 3 right: 4 total: 6 result: 1
//{4} left: 4 right: 4 total: 4 result: 1
//{0} left: 5 right: 4 total: 0 result: 1
//{3} left: 5 right: 5 total: 3 result: 1
//{} left:6 right: 5 total:0 result:1











