package main

import (
	"fmt"
)

func plus(a , b int)int{
	return a+b
}

func minus(a , b int)int{
	var big , small int

	if a>b {
		big , small = a , b 
	}else{
		big , small = b , a
	}

	return big-small
}

func vals()(int , int){
	return 3,7
}

func variadic(nums ...int)int{
	sum := 0
	for _,num := range nums {
		sum+=num
	}
	return sum

}

func factorial(num int)int{
	if(num == 0){
		return 1
	}
	return num*factorial(num-1)
}

func main(){

	nums := []int {1,2,5,4,9,7,8}

	fmt.Println(minus(5,10))
	fmt.Println(plus(5,10))
	
	a , b := vals()
	fmt.Println(variadic(1,2,3,4,5,6,7,8,9))
	fmt.Println(variadic(nums...))

	fmt.Println(a,b)
	fmt.Println(factorial(7))

}