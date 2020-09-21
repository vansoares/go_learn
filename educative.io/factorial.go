package main

import "fmt"

func main() {
	fac := Factorial(3)

	fmt.Println(fac)
}

func Factorial(n uint64) (fac uint64) {

	if n <= 1 {
		return 1
	}
	fac = n * Factorial(n-1)
	return
}
