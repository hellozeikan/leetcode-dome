package main

import "fmt"

func main() {
	for i := 0; i < 2; i++ {
	A:
		for j := 0; j < 10; j++ {
			if j%2 == 0 {
				continue A
			}
			fmt.Println(i, j)
		}
	}
}
func findCircleNum(isConnected [][]int) (res int) {
	n := len(isConnected)
	vis := make([]bool, n)
	var dfs func(n int)
	dfs = func(i int) {
		vis[i] = true
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 && !vis[j] {
				dfs(j)
			}
		}
	}
	for i := 0; i < n; i++ {
		if !vis[i] {
			res++
			dfs(i)
		}
	}
	return
}
