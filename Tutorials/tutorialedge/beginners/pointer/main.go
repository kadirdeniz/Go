package main

import "fmt"

func myTest(a int) {
	a += 3
	fmt.Println(a)
}

func myPointerFunc(a *int) {
	*a += 3
	fmt.Println(*a)
}

func main() {
	a := 2

	myTest(a)      //5
	fmt.Println(a) //2

	myPointerFunc(&a) //5

	myTest(a)      //8
	fmt.Println(a) //5

	var age *int
	age = new(int)
	*age = 20
	fmt.Println(age)
	fmt.Println(&age)

}
