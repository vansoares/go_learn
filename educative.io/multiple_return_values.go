package main

import "fmt"

func main() {
	s, p, d := sumProductDiff(3, 4)
	fmt.Println(s, p, d)
}

func sumProductDiff(i, j int) (s, p, d int) {

	return i + j, i * j, i - j
}
