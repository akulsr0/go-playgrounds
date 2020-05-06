package main

import (
	"fmt"
)

func main() {
	fmt.Print("Welcome to Calculator")
	fmt.Println("=========================>")
	fmt.Println("Enter two numbers: ")
	var a, b, choice int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Println("Enter Choice: ")
	fmt.Println("1. Add 2. Subtract 3. Multiply 4. Divide")
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1: {
		fmt.Println(add(a,b))
		break
	}
	case 2: {
		fmt.Println(subtract(a,b))
		break
	}
	case 3:{
		fmt.Println(multiply(a,b))
		break
	}
	case 4: {
		fmt.Println(divide(a,b))
		break
	}
	default:
		fmt.Println("Invalid Choice!")
		break
	}
}
