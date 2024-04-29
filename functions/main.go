package main

import (
	"errors"
	"fmt"
)

// Go is a functional-first language. It means that go prefers
// the functional paradigm over others but does not enforce it.

// The basic structure of a functions is like below
// func <functionName>(argName <typeArg>, argsName <typeArgs>) return type(s)

func allOperations(n1 int, n2 int) (int, int, int, float64, int) {

	// **Just a little parentesis
	// now you have some examples of Arithmetics operators

	// It's important to say that you cannot use any operator with variables of different types
	// you can't sum, subtract, compare, you can do nothing
	// to work with different types you must firs convert the
	// values

	sum := n1 + n2
	subtraction := n1 - n2
	multiplication := n1 * n2
	division := float64(n1 / n2)
	mod := n1 % n2
	return sum, subtraction, multiplication, division, mod
}

func main() {
	var num1 int = 2
	var num2 int = 9

	// A funcion can have many arguments and many returns

	rsum, rsubtraction, rmultiplication, rdivision, rmod := allOperations(num1, num2)
	fmt.Println(num1, num2, rsum, rsubtraction, rmultiplication, rdivision)

	// A function is a type, and can be assigned to a variable
	// For that reason a function can return another function
	// as you can be seen below.

	// In time to mention, If you want to discard one
	// of the returns, you can assign it to _
	// as can be seen below

	functionB, _ := getFuctionFor("sum")

	// See next the example os error return,
	// that was ignored above but it is important
	// now

	functionB, error1 := getFuctionFor("invalid")
	if error1 != nil {
		fmt.Println(error1)
	}

	fmt.Println(functionB(num1, num2))

	fmt.Println(num1, num2, rsum, rsubtraction, rmultiplication, rdivision, rmod)
}

// This is an example of function return functions
// I'm not sure how good it is, as you can see, it may get quite confuse

func getFuctionFor(operationName string) (func(n1, n2 int) float32, error) {

	//It can return existing functions
	if operationName == "sum" {
		return sum, nil
	}

	if operationName == "sub" {
		return subtraction, nil
	}

	if operationName == "mult" {
		return multiplication, nil
	}

	if operationName == "div" {
		return division, nil
	}

	if operationName == "mod" {
		return modulus, nil
	}

	// as well as internal functions
	return func(n1, n2 int) float32 {
		return 0
	}, errors.New("invalid operation")
}

// Se below examples of assigning functions to variables previously used

var sum = func(num1 int, num2 int) float32 {
	fmt.Println("sum")
	return (float32)(num1 + num2)
}

var subtraction = func(num1 int, num2 int) float32 {
	fmt.Println("subtraction")
	return (float32)(num1 - num2)
}

var multiplication = func(num1 int, num2 int) float32 {
	fmt.Println("multiplication")
	return (float32)(num1 * num2)
}

var division = func(num1 int, num2 int) float32 {
	fmt.Println("division")
	return (float32)(num1 / num2)
}

var modulus = func(num1 int, num2 int) float32 {
	fmt.Println("modulus")
	return (float32)(num1 % num2)
}

// For now it is enough with function
//  -> for further information about the types of functions
//  -> refer to
//
// By now you are famyliar with parameter and
// might be thinking:
// How are the values passed to the
// functions, by value or by reference?
// It is better explained at pointers/main.go
