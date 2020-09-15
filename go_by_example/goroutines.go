package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	func(msg string) {
		for i := 0; i < 5; i++ {
			fmt.Println(msg)
		}
	}("going")

	go f("goroutine")

	func(msg string) {
		for i := 0; i < 5; i++ {
			fmt.Println(msg)
		}
	}("kkkkkk")

	time.Sleep(time.Second)
	fmt.Println("done")
}
