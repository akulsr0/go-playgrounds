package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Golang Math Package")

	// Ceil
	// Ceil returns the least integer value greater than or equal to x.
	fmt.Println(math.Ceil(12.2)) // 13

	// Floor
	// Floor returns the greatest integer value less than or equal to x.
	fmt.Println(math.Floor(12.2)) // 12

	// Pow
	// Pow returns x**y, the base-x exponential of y.
	fmt.Println(math.Pow(2,4)) // 16

	// Mod
	// Mod returns the floating-point remainder of x/y. The magnitude of the result is less than y and its sign agrees with that of x.
	fmt.Println(math.Mod(7,4)) // 3

	// Pi
	// 3.14159265358979323846264338327950288419716939937510582097494459
	fmt.Println(math.Pi) // 3.141592653589793

	// Sin
	// Sin returns the sine of the radian argument x.
	fmt.Println(math.Sin(math.Pi/2)) // 1

	// Cos
	// Cos returns the cosine of the radian argument x.
	fmt.Println(math.Cos(0)) // 1

	// Tan
	// Tan returns the tangent of the radian argument x.
	fmt.Println(math.Tan(math.Pi/4)) // 1

	// Log
	// Log returns the natural logarithm of x.
	fmt.Println(math.Log(1)) // 0

	// Log10
	// Log10 returns the decimal logarithm of x. The special cases are the same as for Log.
	fmt.Println(math.Log10(100)) // 2

	// Log2
	// Log2 returns the binary logarithm of x. The special cases are the same as for Log.
	fmt.Println(math.Log2(256)) // 8

	// Max
	// Max returns the larger of x or y.
	fmt.Println(math.Max(10,20)) // 20

	// Sqrt
	// Sqrt returns the square root of x.
	fmt.Println(math.Sqrt(144))

}
