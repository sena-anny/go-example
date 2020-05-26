package main

import "fmt"

func main() {
	// var a [2]string
	// a[0] = "Hello"
	// a[1] = "World"
	// fmt.Println(a[0], a[1])
	// fmt.Println(a)

	// primes := [6]int{2, 3, 4, 5, 6, 7}
	// var s []int = primes[1:4]
	// fmt.Println(s)

	// names := [4]string{
	// 	"John",
	// 	"Paul",
	// 	"George",
	// 	"Ringo",
	// }
	// fmt.Println(names)

	// a := names[0:2]
	// b := names[1:3]
	// fmt.Println(a, b)

	// b[0] = "XXX"
	// fmt.Println(a, b)
	// fmt.Println(names)

	// q := []int{2, 3, 5, 7, 11, 13}
	// fmt.Println(q)

	// r := []bool{true, false, true, true, false, true}
	// fmt.Println(r)

	// s := []struct {
	// 	i int
	// 	b bool
	// }{
	// 	{2, true},
	// 	{3, false},
	// 	{5, true},
	// 	{7, true},
	// 	{11, false},
	// 	{13, true},
	// }
	// fmt.Println(s)

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	var l []int
	fmt.Println(l, len(l), cap(l))
	if l == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
