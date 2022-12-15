package main

import (
	"fmt"
	"math"
)

const INF = 0x3f3f3f3f

func dijkstra(graph [][]int, start int) []int {
	n := len(graph)         // 图中顶点个数
	visit := make([]int, n) // 标记已经作为中间结点完成访问的顶点
	dist := make([]int, n)  // 存储从前点到其他顶点的最短路径
	for i := 0; i < n; i++ {
		dist[i] = graph[start][i] // 初始化遍历起点
	}
	visit[start] = 1 // 标记初始顶点
	var minDist, midNode int
	// 更新其他顶点最短路径，循环n次
	for i := 0; i < n; i++ {
		minDist = INF // 存储从起点到其他未被访问的结点中的最短路径
		midNode = 0   // 存储最短路径的结点编号
		// 遍历n个顶点，寻找未被访问且距离为起始位置到该点距离最小的顶点
		for j := 0; j < n; j++ {
			if visit[j] == 0 && minDist > dist[j] {
				minDist = dist[j] // 更新未被访问结点的最短路径
				midNode = j       // 更新顶点编号
			}
		}
		// 以midNode为中间结点，再循环遍历其他节点更新最短路径
		for j := 0; j < n; j++ {
			// 若该节点未被访问且找到更短路径即更新最短路径
			if visit[j] == 0 && dist[j] > dist[midNode]+graph[midNode][j] {
				dist[j] = dist[midNode] + graph[midNode][j]
			}
		}
		visit[midNode] = 1 // 标记已访问
	}
	return dist
}
func networkDelayTime(times [][]int, n int, k int) int {
	arr := make([][]int, n)
	for j := 0; j < n; j++ {
		a := make([]int, n)
		for i := 0; i < n; i++ {
			a[i] = INF
		}
		arr[j] = a
	}
	for _, val := range times {
		arr[val[0]-1][val[1]-1] = val[2]
		arr[val[1]-1][val[0]-1] = val[2]
	}
	dist := dijkstra(arr, k)
	max := 0
	for _, v := range dist {
		if max <= v {
			max = v
		}
	}
	return max
}
func main() {
	// 带权值邻接矩阵
	// var gp = [][]int{
	// 	{0, 100, 1200, INF, INF, INF},
	// 	{100, 0, 900, 300, INF, INF},
	// 	{1200, 900, 0, 400, 500, INF},
	// 	{INF, 300, 400, 0, 1300, 1400},
	// 	{INF, INF, 500, 1300, 0, 1500},
	// 	{INF, INF, INF, 1400, 1500, 0},
	// }
	// dist := dijkstra(gp, 1)
	// fmt.Println(dist)
	fmt.Println(checkPowersOfThree(21))
}

func checkPowersOfThree(n int) bool {
	arr := []int{}
	for i := 0; i <= 17; i++ {
		arr = append(arr, int(math.Pow(3, float64(i))))
	}
	find := func() int {
		if n < arr[0] {
			return -1
		}
		for i := 0; i <= len(arr); i++ {
			if n > arr[i] {
				continue
			}
			if n == arr[i] {
				return n
			}
			k := i - 1
			if i == 0 {
				k++
			}
			temp := arr[k]
			arr = arr[:k+copy(arr[k:], arr[k+1:])]
			return temp
		}
		return 0
	}
	for n != 0 {
		t := find()
		if t == -1 {
			return false
		}
		n -= t
	}
	return true
}
func checkPowersOfThree1(n int) bool {
	arr := []int{}
	for i := 0; i <= 17; i++ {
		arr = append(arr, int(math.Pow(3, float64(i))))
	}
	for i := 17; i >= 0; i-- {
		if n >= arr[i] {
			n -= arr[i]
		}
	}
	return n == 0
}
