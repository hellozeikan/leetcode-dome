package main

func main() {
	pivotInteger(8)
}
func pivotInteger(n int) int {
	if n == 1 {
		return 1
	}
	pix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		pix[i] = pix[i-1] + i
	}
	for i := n - 1; i > 0; i-- {
		if pix[i] < pix[n]-pix[i] {
			return -1
		}
		if pix[i] == pix[n]-pix[i]+i {
			return i
		}
	}
	return -1
}
func appendCharacters(s string, t string) int {
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			j++
		}
		i++
	}
	return m - j
}
