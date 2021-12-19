package main

import "fmt"

func bubleSort(array []int) []int {
	for x := 0; x < len(array)-1; x++ {
		for y := x; y < len(array)-1; y++ {
			if array[y] > array[y+1] {
				temp := array[y+1]
				array[y] = array[y+1]
				array[y+1] = temp
			}
		}
	}
	return array
}

func main() {
	arr := [...]int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	fmt.Println(bubleSort(arr[:]))
}
