package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Println("Enter two numbers x and y: ")
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)

	// Simple if
	if x==y {
		fmt.Println("x == y")
	}

	// Simple if-else
	if x%2==0 {
		fmt.Println("x is even.")
	} else {
		fmt.Println("x is odd")
	}

	// if condition with an optional statement
	if i:=x+y; i%2==0 {
		fmt.Println("x+y is even")
	}

	// Else-if ladder
	if i:=x+y; i<10 {
		fmt.Println("x+y is less than 10")
	} else if i>=10 && i<=20 {
		fmt.Println("x+y is between 10 and 20")
	} else {
		fmt.Println("x is greater than 20")
	}

	// Switch Case
	switch sum:=x+y; sum {
	case 10:
		fmt.Println("Sum is 10")
	case 11:
		fmt.Println("Sum is 11")
	case 12:
		fmt.Println("Sum is 12")
	case 13:
		fmt.Println("Sum is 13")
	case 14:
		fmt.Println("Sum is 14")
	case 15:
		fmt.Println("Sum is 15")
	default:
		fmt.Println("Default Condition")
	}
	
}
