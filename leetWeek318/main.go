package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(totalCost([]int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}, 11, 2))
}
func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			continue
		}
		nums[i] *= 2
		nums[i+1] = 0
	}
	arr := make([]int, len(nums))
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			arr[j] = nums[i]
			j++
		}
	}
	return nums
}
func maximumSubarraySum(nums []int, k int) int64 {
	ans, sum := 0, 0
	cnt := map[int]int{}
	for _, x := range nums[:k-1] {
		cnt[x]++
		sum += x
	}
	for i := k - 1; i < len(nums); i++ {
		cnt[nums[i]]++ // 移入元素
		sum += nums[i]
		if len(cnt) == k && sum > ans {
			ans = sum
		}
		x := nums[i+1-k]
		cnt[x]-- // 移出元素
		if cnt[x] == 0 {
			delete(cnt, x) // 重要：及时移除个数为 0 的数据
		}
		sum -= x
	}
	return int64(ans)
}
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func totalCost(costs []int, k int, candidates int) int64 {
	sort.Ints(costs)
	fmt.Println(costs)
	res := 0
	for i := 0; i < len(costs); i++ {
		if len(costs)-i < candidates && i < len(costs) && k > 0 {
			fmt.Println("xxx... ", costs[i])
			res += costs[i]
			break
		}
		if k > 0 {
			fmt.Println(costs[i])
			res += costs[i]
			k--
		}
	}
	return int64(res)
}
