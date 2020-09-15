package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Vanessa", age: 26})
	fmt.Println(&person{name: "Maria Paula", age: 24})

	mamate := newPerson("Thomás")
	mamate.age = 8

	fmt.Println(mamate)

	s := &person{name: "Fernanda", age: 26}
	age := s.age
	fmt.Println(age)
}
