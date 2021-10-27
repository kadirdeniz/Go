package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main(){
	log.SetPrefix("greetings: ")
    log.SetFlags(0)

	ingilizce , err := greetings.Merhaba("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ingilizce)
}