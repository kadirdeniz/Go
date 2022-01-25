package main

import (
	"fmt"
	"sync"
)

var lock = sync.Mutex{}

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("Creating single instance now.")
		singleInstance = &single{}
		return singleInstance
	}
	fmt.Println("Single instance already created.")
	return singleInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	fmt.Scanln()
}
