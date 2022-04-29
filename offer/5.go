package main

import "fmt"

// 剑指 Offer II 005. 单词长度的最大乘积

//输入: words = ["abcw","baz","foo","bar","fxyz","abcdef"]
//输出: 16
//解释: 这两个单词为 "abcw", "fxyz"。它们不包含相同字符，且长度的乘积最大。

// 思路：bitmap + 与运算

func main() {
	fmt.Println(maxProduct([]string{"abcw","baz","foo","bar","fxyz","abcdef"}))
}

func maxProduct(words []string) int {
	r := 0
	m := make([]int, len(words))
	for j, w := range words {
		for _, _w := range w {
			m[j] |= 1 << (_w - 'a')
		}
	}
	for i := 0;i < len(words);i++ {
		for j := i + 1;j < len(words);j++ {
			if m[i] & m[j] == 0 {
				if len(words[i]) * len(words[j]) > r {
					r = len(words[i]) * len(words[j])
				}
			}
		}
	}
	return r
}