package main

import "fmt"

var mode []int = []int{}

func ArrayChallenge(arr []int) int {

	for x := 0; x <= len(arr)-1; x++ {
		val := arr[x]
		for y := 0; y <= len(arr)-1; y++ {
			if x == y {
				continue
			}
			if arr[y] == val {
				mode = append(mode, arr[y])
			}
		}
	}
	if len(mode) == 0 {
		return -1
	}
	return mode[0]
}

func main() {
	fmt.Println(ArrayChallenge(readline()))
}
