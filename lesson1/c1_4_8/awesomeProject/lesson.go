package main

import "fmt"

func main() {
	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)
	fmt.Println(board[1])
	fmt.Println(board[1][2])
}
