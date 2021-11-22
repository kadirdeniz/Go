package main

import "fmt"

// Integers
// all numeric types default to 0

// unsigned int with 8 bits
// Can store: 0 to 255
// var myint uint8

// signed int with 8 bits
// Can store: -127 to 127
// var myint2 int8

func main() {

	fmt.Println("Hello World")

	var myint int8
	for i := 0; i <= 129; i++ {
		myint++
	}
	fmt.Println(myint)

	// Conversion

	var men int8
	men = 5

	var women int16
	women = 6

	var people int

	// people = men+women

	people = int(men) + int(women)
	fmt.Println(people)

	var name string = "Kadir"
	fmt.Println(name)
}
