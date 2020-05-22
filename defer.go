package main

import "fmt"

func main() {
	fmt.Println("conunting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	defer fmt.Println("world")
	fmt.Println("hello")
	fmt.Println("done")
}
