package main

import "fmt"

//剑指 Offer II 010. 和为 k 的子数组
//给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数。

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1, 1}, 3))
}

// [1, 1, 1, 1] : 3

// 1 1
// 1 1 2
// 1 1 1 3
//[1, 1, 1, 1, 1]

func subarraySum(nums []int, k int) int {
	preSum := 0
	count := 0
	m := make(map[int]int)
	m[0] = 1
	for i := 0;i < len(nums);i++ {
		preSum += nums[i]
		if _, ok := m[preSum - k];ok {
			fmt.Println(m[preSum - k])
			count += m[preSum - k]
		}
		m[preSum]++
	}
	return count
}

// 前缀和：元素i的前缀和等于 nums[0] + nums[1] + ... + nums[i]
// 定义前缀和数组preSum
// 某个连续子数组[i, j]的和等于k，可以理解为k = preSum[j] - preSum[i - 1]
// 比如对于数组：[1, 2, 3, 4], 第二个元素到第四个元素的和(2+3+4) 等价于 preSum[4] - preSum[2-1] = (1+2+3+4)-(1)
// 由k = preSum[j] - preSum[i - 1] 可得 preSum[i - 1] = preSum[j] - k
//也就是说 preSum[i]的个数就是满足连续子数组和=k的个数

