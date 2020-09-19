package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 10
	m["k2"] = 90

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	delete(m, "k2")
	fmt.Println("map after delete:", m)

	teste, prs := m["k1"]
	fmt.Println("prs:", prs)
	fmt.Println("teste:", teste)

	n := map[string]int{"vanessa": 1, "santos": 2, "soares": 3}
	fmt.Println("map:", n)
}