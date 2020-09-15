package main

import "fmt"

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res2, res3 := multipleReturns(3, 4)
	fmt.Println(res2, res3)

	sum(1, 2, 3, 4)
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	sum(nums[2:5]...)
}

func plus(a int, b int) int {
	return a + b
}

func multipleReturns(a, b int) (int, int) {
	return a + b, a - b
}

//variadic functions
func sum(nums ...int) {
	fmt.Println(nums, " ")
}
