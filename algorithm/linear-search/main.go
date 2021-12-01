package main

import "fmt"

func linearSearch(arr []int, givenValue int) (bool, int) {

	var isExists bool = false
	var numberOf int = 0

	for x := 0; x < len(arr); x++ {
		if arr[x] == givenValue {
			if !isExists {
				isExists = true
			}
			numberOf++
		}
	}
	return isExists, numberOf
}

func main() {
	arr := [...]int{1, 2, 3, 3, 5, 6, 7, 8, 9}
	isExists, numberOf := linearSearch(arr[:], 3)

	fmt.Println("Is Number Exists = ", isExists)
	fmt.Println("How Many Given Value Exists in Array = ", numberOf)
}
