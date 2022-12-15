package main

import (
	"fmt"
	"math"
	"sort"
)

// n数之和
var res [][]int
var target int

func main() {
	target = 15
	nSum([]int{1, 3, 1, 2, 7, 5, 4, 8}, 5)
	fmt.Println(res)
}
func nSum(nums []int, N int) {
	sort.Ints(nums)
	arr := []int{}
	help(0, N, arr, nums)
}
func help(ind, N int, arr, nums []int) {
	l := len(nums)
	if ind == l || N < 3 {
		return
	}
	for i := ind; i < l; i++ {
		if i > ind && nums[i] == nums[i-1] {
			continue
		}
		if N > 3 {
			help(i+1, N-1, append(arr, nums[i]), nums)
		}
		left := i + 1
		right := l - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right] + numsum(arr)
			if sum == target {
				r := append(arr, nums[i], nums[left], nums[right])
				if len(r) == 5 {
					res = append(res, r)
				}
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
}
func numsum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

// 位图存储bool结果, 1为true
func weitu() {
	a := 4
	a |= a << 1
	b := math.Pow(2, 2)
	fmt.Println(int(b) & a)
	fmt.Printf("%b", a)
}
