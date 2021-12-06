package main

import "fmt"

var name string

func init() {
	fmt.Println("Init Function Invoked")
	name = "Kadir"
}

func init() {
	fmt.Println("This is Init Function Too")
	name = "Yunus"
}

func init() {
	fmt.Println("This is Init Function Too")
	name = "Emre"
}

func main() {
	fmt.Println("Server Started")
	fmt.Println("Name :" + name)
}
