package main

import (
	"fmt"
	"strconv"
)

func main() {
	var name string // variable declaration
	name = "Akul Srivastava" // variable assignment

	var email string = "akulsr0@gmail.com" // variable initialization
	// Above line says type can be omitted, which means when initializing a variable
	// with some value. Go can infer the data type of that variable.

	// Type Inference
	var age = 21 // age is int
	var isSuperUser = true // isSuperUser is bool

	// Multiple variable declaration
	var x, y int = 10, 20

	// Declaring variables of different type in single statement
	var (
		aInteger = 11
		aFloat = 12.424
		aString = "Hello Golang"
		aBool = false
	)

	// ShortHand Declaration (No need of keyword var)
	shortHand := "This is initialized using shorthand expression."

	// Multiple ShortHand Declaration
	width, height := 15, 30

	fmt.Println("Name: ", name)
	fmt.Println("Email: ", email)
	fmt.Printf("Age is %d of type %T \n", age, age)
	fmt.Printf("isSuperUser is %s of type %T\n", strconv.FormatBool(isSuperUser), isSuperUser)
	fmt.Println("x: ", x, "and y: ", y)
	fmt.Println("aInteger: ", aInteger)
	fmt.Println("aFloat: ", aFloat)
	fmt.Println("aString: ", aString)
	fmt.Println("aBool: ", aBool)
	fmt.Println("shortHand: ", shortHand)
	fmt.Println("width: ", width, "height: ", height)
}