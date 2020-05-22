package main

import (
	"fmt"
	"math"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

// func Sqrt(x float64) float64 {

// }

func main() {
	fmt.Println("Go runs on")
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("OS X.")
	// case "linux":
	// 	fmt.Println("Linux.")
	// default:
	// 	fmt.Printf("%s.", os)
	// }
	// today := time.Now().Weekday()
	// fmt.Println(today + 1)
	// switch time.Saturday {
	// case today + 0:
	// 	fmt.Println("Today.")
	// case today + 1:
	// 	fmt.Println("Tommorow is", today+1)
	// case today + 2:
	// 	fmt.Println("In two days.")
	// default:
	// 	fmt.Println("Too far away.")
	// }
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening", t.Hour())
	}
	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )
	// fmt.Println(sqrt(2), sqrt(-4))
	// sum := 1
	// for sum < 100 {
	// 	sum += sum
	// 	fmt.Println(sum)
	// }
	// fmt.Println(sum)
}
