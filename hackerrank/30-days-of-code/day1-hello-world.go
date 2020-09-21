package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	msg, _ := reader.ReadString('\n')

	fmt.Println("Hello, World.")
	fmt.Println(msg)
}
