package main

import (
	"fmt"
)

func main(){
	nums := []int{2,3,4}
	sum := 0

	for _,num := range nums {
		sum +=num
	}

	fmt.Println(sum)


	for i , z := range nums {
		fmt.Println(fmt.Sprintf("index => %v , value => %v",i,z))
	}

	kvs := map[string]string{"a":"apple","b":"banana"}
	for i,v := range kvs {
		fmt.Println(fmt.Sprintf("index => %v , value => %v",i,v))
	}

	for v := range kvs {
		fmt.Println(fmt.Sprintf("index => %v",v))

	}
}