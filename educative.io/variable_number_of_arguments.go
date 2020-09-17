package main

import "fmt"

func main() {
	total := sumInts(1, 2, 3, 4, 5)

	fmt.Println(total)
}

func sumInts(list ...int) (sum int) {
	var total int
	for index, num := range list {
		total += num
		fmt.Println(index)
	}
	return total
}
