package main

import (
	"fmt"
	"math"
)

// "math/cmplx"

// var c, python, java bool
// var i, j int = 1, 2
// var (
// 	ToBe   bool       = false
// 	MaxInt uint64     = 1<<64 - 1
// 	z      complex128 = cmplx.Sqrt(-5 + 12i)
// )

const Pi = 3

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	v := 42.34 + 3i // change me!
	fmt.Printf("v is of type %T\n", v)
	// var i int
	// k := 3
	// var c, python, java = true, false, "no"
	// fmt.Printf("type: %T value: %v\n", ToBe, ToBe)
	// fmt.Printf("type: %T value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("type: %T value: %v\n", z, z)
	// var i int
	// var f float64
	// var b bool
	// var s string
	// fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
