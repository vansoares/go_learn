package main

import "fmt"

func main() {
	//Make = create empty slice with non zero length
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")

	//k := []string{"g", "h"}

	s = append(s, "a", "k")

	fmt.Println("append 2 arrays: ", s)

	l := s[2:5]
	fmt.Println("slice:", l)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("TWO_D:", twoD)
}