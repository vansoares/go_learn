package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for index, num := range nums {
		sum += num
		fmt.Println("-", index)
		fmt.Println(">", num)
	}

	kvs := map[string]string{"a": "apple", "b": "bananas", "c": "caju"}
	for key, value := range kvs {
		fmt.Printf("%s -> %s \n", key, value)
	}
}
