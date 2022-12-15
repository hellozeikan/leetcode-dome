package main

import "fmt"

func main() {
	fmt.Println(isIdealPermutation([]int{0, 1, 2, 5, 4, 3}))
}
func isIdealPermutation(nums []int) bool {
	l := len(nums)
	// 比如
	// max 从len(nums) -1 往前推的最小值
	// nums[i] > nums[i+1] && nums[i] > max[i+2return false
	maxArr := make([]int, l)
	maxArr[l-1] = nums[l-1]
	for i := l - 2; i >= 0; i-- {
		if nums[i] < maxArr[i+1] {
			maxArr[i] = nums[i]
		} else {
			maxArr[i] = maxArr[i+1]
		}
	}
	fmt.Println(maxArr)
	fmt.Println(nums)
	for i := 0; i < l-2; i++ {
		if nums[i] == 5 {
			fmt.Println(nums[i], nums[i+1], maxArr[i+2])
		}
		if nums[i] > nums[i+1] && nums[i] > maxArr[i+2] {
			return false
		}
	}
	//  method 1
	for i := 0; i < len(nums); i++ {
		if abs(nums[i]-i, 0) >= 2 {
			return false
		}
	}
	//  method 2
	max := nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i] < max {
			return false
		}
		if max < nums[i-1] {
			max = nums[i-1]
		}
	}
	return true
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
