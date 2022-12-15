package main

import (
	"fmt"
	"time"
)

func main() {
	// os.Mkdir("test", os.ModePerm)
	// fmt.Println(gcd(6, 3))
	// fmt.Println(numSubarrayBoundedMax([]int{2, 9, 2, 5, 6}, 2, 8))
	// expressiveWords("heeellooo", nil)
	t := time.Now()
	dd, _ := time.ParseDuration("24h")
	t.Add(7 * dd).Format("2006-01-02")
	fmt.Println()
}
func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	k := 0
	for target > 0 {
		k++
		target -= k
	}
	if target == 0 {
		return k
	}
	return k + 1 + k%2
}
func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}

/*
给你一个整数数组 nums 和两个整数：left 及 right 。找出 nums 中连续、非空且其中最大元素在范围 [left, right] 内的子数组，并返回满足条件的子数组的个数。

生成的测试用例保证结果符合 32-bit 整数范围。
*/
func numSubarrayBoundedMax(nums []int, left int, right int) int {
	res := 0
	l := len(nums)
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = checkone(nums[i], left, right)
	}
	for i := 0; i < l; i++ {
		if arr[i] == -1 {
			continue
		}
		for j := i; j < l; j++ {
			if arr[j] == -1 {
				break
			}
			if check(arr[i : j+1]) {
				res++
			}
		}
	}
	return res
}

// 范围内1， 小于0， 大于为-1
func checkone(a, b, c int) int {
	if b <= a && a <= c {
		return 1
	}
	if a < b {
		return 0
	}
	return -1
}
func check(arr []int) bool {
	for _, val := range arr {
		if val == 1 {
			return true
		}
	}
	return false
}

func expressiveWords(s string, words []string) int {
	res := 0
	for _, val := range words {
		if help(s, val) {
			res++
		}
	}
	return res
}
func help(t, s string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] != t[j] {
			return false
		}
		ch := s[i]
		cntI := 0
		for i < n && s[i] == ch {
			cntI++
			i++
		}
		cntJ := 0
		for j < m && t[j] == ch {
			cntJ++
			j++
		}
		if cntI < cntJ || cntI > cntJ && cntI < 3 {
			return false
		}
	}
	return i == n && j == m
}
