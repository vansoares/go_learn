package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	c4 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c3 <- "3"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c4 <- "4"
	}()

	for i := 0; i < 4; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case msg3 := <-c3:
			fmt.Println("received", msg3)
		case msg4 := <-c4:
			fmt.Println("received", msg4)
		}
	}
}
