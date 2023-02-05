package ch3

import "fmt"

func SimpleGoroutine() {
	go sayHello()
}

func sayHello() {
	fmt.Println("Hello")
}
