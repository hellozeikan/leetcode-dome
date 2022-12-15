package main

import (
	"fmt"
	"sort"
)

func main() {
	largestSumOfAverages([]int{9, 1, 2, 3, 9}, 3)
}
func largestSumOfAverages(nums []int, k int) float64 {
	var res float64 = 0
	sort.Ints(nums)
	l := len(nums)
	sum := 0
	for i := 0; i < l; i++ {
		if i > l-k {
			res += float64(nums[i])
		} else {
			sum += nums[i]
		}
	}
	fmt.Println(res, sum)
	res += float64(sum) / float64(l-k)
	return res
}
