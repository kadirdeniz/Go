package main

import "fmt"

func traverse(arr []int) []int {
	var response []int
	for x := len(arr) - 1; x > 0; x-- {
		response = append(response, x)
	}
	return response
}

func main() {
	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(traverse(arr))
}
