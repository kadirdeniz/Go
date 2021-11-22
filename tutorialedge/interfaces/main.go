package main

import "fmt"

type Engineer struct {
	Name string
}

type Manager struct {
	Name string
}

func (e *Engineer) getName() string {
	return e.Name
}

func (m *Manager) getName() string {
	return m.Name
}

type Employee interface {
	getName() string
}

func printDetails(e Employee) {
	fmt.Println(e.getName())
}

func main() {

	var engineer Engineer = Engineer{
		Name: "Kadir",
	}

	var manager Manager = Manager{
		Name: "Yunus Emre",
	}

	printDetails(&engineer)
	printDetails(&manager)

}
