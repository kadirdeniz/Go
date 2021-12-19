package main

import (
	"family/father"
	"family/father/son"
	"fmt"
)

func main() {
	f := new(father.Father)
	fmt.Println(f.Data("Mr. Jeremy Maclin "))

	s := new(son.Son)
	fmt.Println((s.Data("Kadir Deniz ")))

}
