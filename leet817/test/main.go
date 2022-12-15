package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof(One{}))
	fmt.Println(unsafe.Sizeof(Two{}))
	// fmt.Println(totalFruit([]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 3, 3, 4, 1}))
}

type One struct {
	s struct{}
	x int
}
type Two struct {
	x int
	s struct{}
}

func totalFruit(fruits []int) int {
	s := []int{}
	for _, elem := range fruits {
		s = append(s, elem)
	}
	len := len(fruits)
	if len == 1 {
		return 1
	}
	res := 0
	i, j := 0, 2
	for i <= len && i <= j && j <= len {
		if j < len && fruits[i] == fruits[j] {
			j++
			continue
		}
		if isTwoValue(s[i:j]) {
			// fmt.Println(s[i:j])
			res = max(res, j-i)
			j++
		} else {
			i++
		}
	}
	return res
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func isTwoValue(arr []int) bool {
	a, b := -1, -1
	for _, v := range arr {
		if a != b && a != v && b != v {
			return false
		}
		if a != v {
			b = v
		}
		if a == -1 {
			a = v
		}

	}
	return true
}



