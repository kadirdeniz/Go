package main

import "fmt"

type Employee struct {
	Name string
}

func (e *Employee) setName(newName string) {
	e.name = newName

}

func (e *Employee) getName() string {
	return e.Name
}

func main() {
	var employee Employee = Employee{
		Name: "Yunus Emre",
	}

	fmt.Println(employee.getName())

	employee.setName("Kadir")

	fmt.Println(employee.getName())
}
