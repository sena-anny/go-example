package main

import (
	"fmt"
	"strings"
)

func main() {
	// primes := []int{2, 3, 5, 7, 11, 13}
	// fmt.Println(primes)

	// r := []bool{true, false, true, true, false}
	// fmt.Println(r)

	// o := []struct {
	// 	i int
	// 	b bool
	// }{
	// 	{2, true},
	// 	{3, false},
	// 	{5, true},
	// 	{7, true},
	// }

	// fmt.Println(o)

	// var s []int = primes[1:4]
	// fmt.Println(s)
	// s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// printSlice(s)

	// s = s[:0]
	// printSlice(s)

	// s = s[:4]
	// printSlice(s)

	// s = s[2:]
	// printSlice(s)

	// s = s[:6]
	// printSlice(s)
	a := make([]int, 5)
	// a = []int{1, 2, 3, 4, 5}
	printSlice("a", a)

	b := make([]int, 0, 5)
	// b = []int{1, 2, 3, 4, 5}
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
