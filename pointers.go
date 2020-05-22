package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	// v := Vertex{1, 2}
	// // p := &v
	// // p.X = 200
	// v.X = 200
	// fmt.Println(v)
	fmt.Println(v1, p.X, v2, v3)

	// i, j := 42, 2701

	// p := &i
	// fmt.Println(i, *p)
	// *p = 21
	// fmt.Println(i, *p)

	// p = &j
	// *p = *p / 37
	// fmt.Println(i, *p, j)
}
