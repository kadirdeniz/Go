package main

import "fmt"

// We’ll be covering the following topics within this tutorial:

// The basics on Function Declaration
// Working with multiple return values

// Function Declaration
func sayMayName(name string, surname string) (string, error) {
	return name + " " + surname, nil
}

func Exported(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fullname, err := fmt.Println(sayMayName("Kadir", "Deniz"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fullname)
}
