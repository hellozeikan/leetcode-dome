package main

func main() {

}

func champagneTower(poured int, query_row int, query_glass int) float64 {
	dp := make([][]float64, query_row+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]float64, query_glass+1)
	}
	dp[0][0] = float64(poured)
	for i := 0; i <= query_row; i++ {
		for j := 0; j <= i; j++ {
			if dp[i][i] < 1 {
				continue
			}
			dp[i+1][j] += (dp[i][j] - 1) / 2
			dp[i+1][j+1] += (dp[i][j] - 1) / 2
		}
	}
	return dp[query_row][query_glass]
}
