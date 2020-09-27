package main

import "fmt"

func main() {
	var array [15]int

	for i := 0; i < 15; i++ {
		array[i] = i
	}

	fmt.Println(array)

}
