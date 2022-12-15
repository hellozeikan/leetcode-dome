package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{1, 2, 3}, {2, 5, 7}, {3, 5, 1}}
	quer := []int{5, 6, 2}
	fmt.Println(maxPoints(arr, quer))
}

func deleteGreatestValue(grid [][]int) int {
	l := len(grid)
	l1 := len(grid[0])
	for i := 0; i < l; i++ {
		sort.Ints(grid[i])
	}
	res := 0
	for i := 0; i < l1; i++ {
		temp := 0
		for j := 0; j < l; j++ {
			if temp < grid[i][j] {
				temp = grid[i][j]
			}
		}
		res += temp
	}
	return res
}

func longestSquareStreak(nums []int) int {
	res := 0
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	for i, _ := range m {
		cnt := 0
		ok := true
		j := i
		for ok {
			cnt++
			j *= j
			_, ok = m[j]
		}
		res = max(res, cnt)
	}
	return res
}

// 2503
func maxPoints(grid [][]int, queries []int) []int {
	res := []int{}
	n := len(grid)
	m := len(grid[0])
	boo := make([][]bool, n)
	for i := 0; i < n; i++ {
		boo[i] = make([]bool, m)
	}
	var dfs func(i, j int, max int) int
	dfs = func(i, j int, max int) (temp int) {

		if i >= n || i < 0 || j >= m || j < 0 {
			return temp
		}
		if boo[i][j] {
			return temp
		}
		if max > grid[i][j] {
			temp++
		} else {
			return temp
		}
		boo[i][j] = true
		return dfs(i-1, j, max) + dfs(i+1, j, max) + dfs(i, j-1, max) + dfs(i, j+1, max) + temp
	}
	for _, v := range queries {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				boo[i][j] = false
			}
		}
		res = append(res, dfs(0, 0, v))
	}
	return res
}

func maxAreaOfIsland(grid [][]int) (res int) {
	n := len(grid)
	m := len(grid[0])
	vis := make([][]bool, n) // 记录走过的点
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, m)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (temp int) {
		if i >= n || i < 0 || j >= m || j < 0 {
			return temp
		}
		if vis[i][j] {
			return temp
		}
		if grid[i][j] == 1 {
			temp++
		} else {
			return temp
		}
		vis[i][j] = true
		return dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1) + temp
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				re := dfs(i, j)
				res = max(re, res)
			}
		}
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
