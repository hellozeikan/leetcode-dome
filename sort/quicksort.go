package main

import "fmt"

func main() {
	// [1,3],[2,2],[3,1]
	arr := [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}
	// fmt.Println(arr)
	// QuickSort1(arr, 0, len(arr)-1)
	fmt.Println(maximumUnits(arr, 10))
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	QuickSort1(boxTypes, 0, len(boxTypes)-1)
	sum := 0
	for i := len(boxTypes) - 1; i >= 0; i-- {
		if truckSize >= boxTypes[i][0] {
			truckSize -= boxTypes[i][0]
			sum += boxTypes[i][0] * boxTypes[i][1]
		} else {
			sum += boxTypes[i][1] * truckSize
			truckSize = 0
		}
	}
	return sum
}
func QuickSort(nums []int, low, high int) {
	if low < high {
		pivo := partition(nums, low, high)
		// 基准不做操作
		QuickSort(nums, low, pivo-1)
		QuickSort(nums, pivo+1, high)
	}
}
func partition(nums []int, low, high int) int {
	pivot := nums[low]
	for low < high {
		for low < high && nums[high] >= pivot {
			high--
		}
		nums[low] = nums[high]
		for low < high && nums[low] <= pivot {
			low++
		}
		nums[high] = nums[low]
	}
	nums[low] = pivot
	return low
}

func QuickSort1(nums [][]int, low, high int) {
	if low < high {
		pivo := partition1(nums, low, high)
		// 基准不做操作
		QuickSort1(nums, low, pivo-1)
		QuickSort1(nums, pivo+1, high)
	}
}
func partition1(nums [][]int, low, high int) int {
	pivot := nums[low]
	for low < high {
		for low < high && nums[high][1] >= pivot[1] {
			high--
		}
		nums[low] = nums[high]
		for low < high && nums[low][1] <= pivot[1] {
			low++
		}
		nums[high] = nums[low]
	}
	nums[low] = pivot
	return low
}
