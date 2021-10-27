package greetings

import (
	"fmt"
	"errors"
)

func Merhaba(name string) (string,error) {
	if name == "" {
		return "" , errors.New("İsim Alanı Doldurulmak Zorunda")
	}
	var message string
	message = fmt.Sprintf("Merhaba, %v Hoş Geldin",name)
	return message,nil
}

func Hello(name string) (string,error) {
	if name == "" {
		return "" , errors.New("Name Field Need To Fill")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message,nil
}