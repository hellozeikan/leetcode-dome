package main

import (
	"fmt"
	"strconv"
)

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	// var peo People = &Student{}
	// think := "bitch"
	// fmt.Println(peo.Speak(think))
	fmt.Println(numDifferentIntegers("leet1234code234"))
}

// import "fmt"

// func main() {
// 	// arr := [][]byte{
// 	// 	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
// 	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
// 	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
// 	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
// 	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
// 	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
// 	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
// 	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
// 	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
// 	// }
// 	// solveSudoku(arr)
// 	// for i := 0; i < 9; i++ {
// 	// 	for j := 0; j < 9; j++ {
// 	// 		fmt.Print(arr[i][j]-48, " ")
// 	// 	}
// 	// 	fmt.Println()
// 	// }
// 	var a int
// 	fmt.Println(&a)
// 	var p *int
// 	p = &a
// 	*p = 20
// 	fmt.Println(a)
// }
// 数读
func solveSudoku(board [][]byte) {
	rowUse := make([][]bool, 9)
	colUse := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rowUse[i] = make([]bool, 10)
		colUse[i] = make([]bool, 10)
	}
	boxUse := make([][][]bool, 3)
	for i := 0; i < 3; i++ {
		boxUse[i] = make([][]bool, 3)
		for j := 0; j < 3; j++ {
			boxUse[i][j] = make([]bool, 10)
		}
	}

	// 初始化数树
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board); col++ {
			num := board[row][col] - 48
			if 1 <= num && num <= 9 {
				rowUse[row][num] = true
				colUse[col][num] = true
				boxUse[row/3][col/3][num] = true
			}
		}
	}
	recusive(board, rowUse, colUse, boxUse, 0, 0)
}

func recusive(board [][]byte, rowUse [][]bool, colUse [][]bool, boxUse [][][]bool, row, col int) bool {
	if col == 9 {
		col = 0
		row++
		if row == 9 {
			return true
		}
	}
	if board[row][col] == '.' {
		for num := 1; num <= 9; num++ {
			can := !(rowUse[row][num] || colUse[col][num] || boxUse[row/3][col/3][num])
			if can {
				rowUse[row][num] = true
				colUse[col][num] = true
				boxUse[row/3][col/3][num] = true
				board[row][col] = byte(num + 48)
				if recusive(board, rowUse, colUse, boxUse, row, col+1) {
					return true
				}
				// 回溯的核心，在于取消上一次的选择
				board[row][col] = '.'
				rowUse[row][num] = false
				colUse[col][num] = false
				boxUse[row/3][col/3][num] = false
			}
		}
	} else {
		return recusive(board, rowUse, colUse, boxUse, row, col+1)
	}
	return false
}

var m map[int]int = make(map[int]int)

// 题目 1805（e）
func numDifferentIntegers(word string) int {
	l := len(word)
	count := 0
	for i := 0; i < l; i++ {
		if word[i] >= '0' && word[i] <= '9' {
			j := i
			for j < l {
				if word[j] >= '0' && word[j] <= '9' && j != l-1 {
					j++
				} else {
					if j == l-1 {
						 if deal(word[i : j+1]) {
							count++
						 }
					} else {
						if deal(word[i:j]) {
							count++
						}
					}
					i = j
					break
				}
			}
		}
	}
	return count
}

func deal(w string) bool {
	val, _ := strconv.Atoi(w)
	if _, ok := m[val]; !ok {
		m[val] = 1
		return true
	}
	return false
}
