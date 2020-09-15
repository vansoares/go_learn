package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	newMessages := make(chan string, 3)

	newMessages <- "buffered"
	newMessages <- "channel"
	newMessages <- "kk"

	fmt.Println(<-newMessages)
	fmt.Println(<-newMessages)

	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

}
