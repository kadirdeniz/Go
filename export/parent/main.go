package main

import (
	"fmt"
	c "parent/child"
)

func main() {

	c.PrintData()
	a := &c.Data

	fmt.Println("Adress of Data :", a)
	fmt.Println("Value of Data :", *a)
	*a++
	c.PrintData()
}
