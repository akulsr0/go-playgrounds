package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Types in Golang
	// Numberic Types:
	// int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint, float32, float64, complex64, complex128, byte, rune
	// String: Collection of characters
	// Boolean: true or false
	// Derived Types: pointer, array, structure, union, functions, interfaces

	var x int = -10
	fmt.Println(x)    // 10

	// var y uint = -10
	// fmt.Println(y)    // Error: -10 overflows uint

	floatNum := 10.56
	fmt.Printf("floatNum is %f and of type %T\n", floatNum, floatNum)

	complexNumber := 5 + 3i
	fmt.Println("complexNumber: ", complexNumber)

	username := "John Doe"
	greetUser := "Welcome " + username

	fmt.Printf("greetUser: %s and type: %T\n", greetUser, greetUser)

	myBool := false
	fmt.Printf("myBool: %s and type %T\n", strconv.FormatBool(myBool), myBool)
	
}
