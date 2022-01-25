package main

import (
	"fmt"
	"strings"
)

var vowels string = "aeiouAEIOU "

func StringChallenge(str string) int {
	for _, value := range str {
		if strings.Contains(vowels, string(value)) {
			str = strings.Replace(str, string(value), "", 1)
		}
	}
	return len(str)
}

func main() {
	fmt.Println(StringChallenge(readline()))
}
