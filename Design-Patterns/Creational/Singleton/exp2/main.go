package main

import (
	"fmt"

	"github.com/design-patterns/file"
)

func main() {

	instance := file.GetInstance()
	fmt.Println(instance.IsFileExists())
	instance2 := file.GetInstance()
	fmt.Println(instance2.IsFolderExists())

	instance.CreateFolder()
	instance.CreateFile()

}
