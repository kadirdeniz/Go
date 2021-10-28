package main

import (
	"fmt"
)


func main(){
	
	s:=make([]string,3)
	s[0]="kadir"
	s[1]="deniz"
	s[2]="22"
	
	s = append(s,"erkek")
	s = append(s,"yazılım mühendisi")

	c := make([]string,len(s))
	copy(c,s)
	fmt.Println(c)

	m:= make(map [string] string)

	m["isim"] = "Kadir"
	m["soyisim"] = "Deniz"
	m["yas"] = "22"

	fmt.Println(m)
	fmt.Println(m["isim"])
	fmt.Println(len(m))
	delete(m,"yas")
	fmt.Println(m)
	fmt.Println(len(m))


	n:= map[string]int{"foo":1,"bar":2}
	n["foo"] = 3
	fmt.Println(n)
}