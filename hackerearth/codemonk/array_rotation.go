package main

import "fmt"

func main() {
	/* T quantidade de casos de teste
	   N quantidades de elementos no array
	   K passos de rotação
	   A array
	*/

	var N, T, K int
	A := make([]int, 0)

	fmt.Scanf("%d", &T)

	for i := 0; i < T; i++ {

		fmt.Scanln(&N)
		fmt.Scanln(&K)

		fmt.Println("K:", K)
		fmt.Println("N:", N)

		for j := 0; j < N; j++ {
			var value int
			fmt.Scanln(&value)
			A = append(A, value)
		}
	}

	resultado := rotation(A, K)
	fmt.Println(resultado)

}

func rotation(A []int, K int) []int {
	first := A[0 : K+1]
	second := A[K+1 : len(A)]

	return append(second, first...)
}
