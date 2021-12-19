package main

import (
	"fmt"
)



func main(){
	var age = 22
	var name , surname string = "Kadir" , "Deniz"
	var gender = true

	message := fmt.Sprintf("Merhaba %v %v : age %v and gender %v",name,surname,age,defineGender(gender))
	fmt.Println(message)
	var i int
	
	i = 10
	fmt.Println(i)
	i = 20
	fmt.Println(i)
}

func defineGender (gender bool) string{
	if gender == true {
		return "Male"
	}else{
		return "Female"
	}
}