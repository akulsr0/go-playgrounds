package main

import "fmt"

func main() {
	// Functions
	sayHello() // Simple Function
	sayHelloWithName("Akul") // Function with param
	fmt.Println(getSum(5,6)) // Function with return value
	multiply, divide := multiplyAndDivide(7,5)
	fmt.Println(multiply, divide) // Function with multiple return value
	fmt.Println(isEven(10)) // Function with named value return
}

// Simple Function
func sayHello() {
	fmt.Println("Hello")
}

// Function with param
func sayHelloWithName(name string) {
	fmt.Println("Hello ", name)
}

// Function with return value
// Note: (x int, y int) is same as (x, y int)
func getSum(x, y int) int {
	return x+y
}

// Function with multiple return value
func multiplyAndDivide(x, y int) (int, float64) {
	return x*y, float64(x)/float64(y)
}

// Return named values
func isEven(n int) (isEven bool) {
	if isEven=false; n%2==0 {
		isEven = true
	}
	return
}