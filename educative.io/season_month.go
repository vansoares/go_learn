package main

import (
	"errors"
	"fmt"
)

func main() {
	month := 12
	season, _ := getSeason(month)
	fmt.Println(season)
}

func getSeason(number int) (string, error) {
	var season string
	var err error
	switch number {
	case 1, 2, 12:
		season = "Winter"
	case 3, 4, 5:
		season = "Spring"
	case 6, 7, 8:
		season = "Summer"
	case 9, 10, 11:
		season = "Autumn"
	default:
		return season, errors.New("Não é um mês válido!!")
	}
	return season, err
}
