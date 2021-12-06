package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

func Merhaba(name string) (string,error) {
	if name == "" {
		return "" , errors.New("İsim Alanı Doldurulmak Zorunda")
	}
	var message string
	message = fmt.Sprintf(randomFormat(),name)
	return message,nil
}

func Hello(name string) (string,error) {
	if name == "" {
		return "" , errors.New("Name Field Need To Fill")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message,nil
}

// init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
	fmt.Println(formats)

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}

func Hellos(names[]string)(map[string]string,error){
	fmt.Println(names)
	messages:=make(map[string]string)
	fmt.Println(messages)
	for _,name := range names {
		message,err := Hello(name)
		if err != nil {
			return nil,err
		}
		messages[name] = message
	}
	return messages,nil
}