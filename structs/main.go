// struct is the a collections of fields and
// closer you can get to a class in go
package main

import "fmt"

// Simple as that
type Stuff struct {
	v1 string
	v2 string
}

// You can also use struct inside structs
// nothing special here
type CompositeStuff struct {
	id     int
	value1 int
	value2 int
	stuff1 Stuff
}

// But there is something that looks like
// inheritance somehow
type SubStuff struct {
	// In this case SubStuff will have all
	// CompositeStuff fields
	CompositeStuff
	value3 string
	value4 string
}

func main() {
	// you can instantiate strucs just like a variable
	var stuff Stuff
	stuff.v1 = "test1"
	stuff.v2 = "test2"

	// You can instantiate strucs like
	// calling a constructor
	// pay atenttion the we are using braces {} and not parentheses ()
	compositeStuff := CompositeStuff{1, 2, 3, stuff}

	// You can instantiate strucs like
	// calling a constructor with named
	// parameters
	stuff2 := CompositeStuff{id: 2, value1: 3, value2: 2, stuff1: stuff}

	stuff2.stuff1.v2 = "qqqqqqq"
	fmt.Println(stuff2)

	// which allows you instantiate structs
	// even if you dont have all fields
	stuff3 := CompositeStuff{value1: 3}
	stuff3.stuff1 = stuff
	stuff3.stuff1.v1 = "test333"

	fmt.Println(stuff)
	fmt.Println(compositeStuff)
	fmt.Println(stuff)

	// An struct that inherits another one
	// is instantiated as it was receiving
	// its parent as a parameter
	subStuff := SubStuff{CompositeStuff: stuff2, value3: "test"}

	fmt.Println(subStuff)
}

// By now you might be thinking on how to organize your
// structs and funcions
//
// To do so you will need to know about packages
// go to ../packages/main.go from here
