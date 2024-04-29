// # Start from here

// # Below is the anatomy of an executable file

// Everything starts on the main package
// It says to go that this package is an executable

package main

// Here you state the packages you are going to use
// One package per line

import (
	"errors"
	"fmt"
)

// The file must contain a function main in it
// The file name is not important as long it is in the
// main package and contains a main function

// Type the following command to run:

// ´´´go run .\main.go #<filename.go>

// You can also type go build to generate an executable

func main() {

	fmt.Println("Here we Go!")

	// ## Declaring variables

	// Go is a strongly typed language
	// A variable can be declared defining its type

	var number1 int = 1

	var (
		string1 string = "jonh"
		number2 int    = 2
	)

	// And declaring with type inference

	number3 := 3
	number4, string2 := 4, "doe"

	// Important to mention, declaring a variable
	// and not using it throws a compilation error
	// Comment one line below to see it fail

	fmt.Println(number1)
	fmt.Println(string1)
	fmt.Println(number2)
	fmt.Println(number3)
	fmt.Println(number4, string2)

	// This is a constant

	const const1 string = "const string"

	fmt.Println(const1)

	// Which you cannot change.
	// Uncomment below to see it failing
	// cannot assign to const1
	// const1 = 3

	// ## Primitive Data Types
	//
	// You can have int8-64 signed and unsigned

	var (
		i8  int8  = -127
		ui8 uint8 = 255

		i16  int16  = -32767
		ui16 uint16 = 65535

		i32  int32  = -2147483647
		ui32 uint32 = 4294967295

		i64  int64  = -9223372036854775807
		ui64 uint64 = 18446744073709551615

		// And there are also the implicit versions of int and uint
		// When its size is not defined, it will use
		// the computer arquitecture to define it.

		vint  int  = -9223372036854775807
		vuint uint = 18446744073709551615
	)

	fmt.Printf("Ints: %v, %v, %v, %v\n", i8, i16, i32, i64)
	fmt.Printf("UInts: %v, %v, %v, %v\n", ui8, ui16, ui32, ui64)
	fmt.Println(vint, vuint)

	// It you use implicit type declaration
	// the variable will assume the best type for the computer architecture

	idnint := i64
	idpint := i64 * (-1)

	fmt.Printf("idnint: %T: %v\n", idnint, idnint)
	fmt.Printf("idpint: %T: %v\n", idpint, idpint)

	// There are some alias for int and uint
	// as you can see below you can declare an
	// int32 as rune and an uint8 as byte

	var number5 rune = 32
	var number6 byte = 8
	fmt.Println(number5)
	fmt.Println(number6)

	// for floating points there are only 32, and 64 bits

	var float1 float32 = 32.32
	var float2 float64 = 64.64
	fmt.Println(float1)
	fmt.Println(float2)

	// When using implicit declaration it will assumve float.
	// we cannot use that type thoug

	float3 := 23.23
	fmt.Printf("float3: %T: %v\n", float3, float3)

	// Just to mention about the strings
	// there isn't char in go.
	// it will be converted to a number

	var string3 string = "this is a string"

	char := 'A'
	fmt.Println(string3)
	fmt.Printf("it should be the char A ut it's %T:%v\n", char, char)

	// No secrets for booleans

	var boolean1 bool = true
	var boolean2 bool = false

	fmt.Println(boolean1, boolean2)

	// 	lets take the opportunity to talk about
	// 	# Operators (Find examples on arithimetic operatos in functions/main.go)
	// 	First the relational operators

	fmt.Println(2 > 1)
	fmt.Println(1 >= 1)
	fmt.Println(1 == 1)
	fmt.Println('A' != 'B')
	fmt.Println(2 <= 2)
	fmt.Println(1 < 2)

	//Logical operatos
	fmt.Println(boolean1 && boolean2)
	fmt.Println(boolean1 || boolean2)
	fmt.Println(!boolean2)

	//Unary operators
	fmt.Println(number1)
	number1++
	fmt.Println(number1)
	number1--
	fmt.Println(number1)
	number1 += 10
	fmt.Println(number1)
	number1 -= 10
	fmt.Println(number1)
	number1 *= 10
	fmt.Println(number1)
	number1 /= 10
	fmt.Println(number1)
	number1 %= 10
	fmt.Println(number1)

	// Unfortunatly (or not) there aren't ternary operators in go

	// Error is also a type.
	// Refer to errors/main.go to learn more

	var error1 error = errors.New("example error")

	fmt.Println(error1)

	// 	All values has its null value
	// 	for strings is ""
	// 	for numbers are 0
	// 	for an error is nil
	//
	//  You can also create your own types
	//  you can learn more about that on
	//  struct/main.go
	//
	// 	Functions are also a type in go
	// 	go to functions/main.go to learn more
	//

}
