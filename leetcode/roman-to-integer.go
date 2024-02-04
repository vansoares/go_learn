package main

import "fmt"

func romanToInt(s string) int {
	if len(s) < 1 || len(s)>15 {
		return 0
	}

	total := 0
	for index, i := range s {
		switch string(i) {
		case "I":
			if index+1 < len(s) && (string(s[index+1]) == "V" || string(s[index+1]) == "X") {
				total -= 1
				continue
			} 
			total += 1
			
		case "V":
			total += 5
		case "X":
			if index+1 < len(s) && (string(s[index+1]) == "L" || string(s[index+1]) == "C") {
				total -= 10
				continue
			} 
			total += 10
		case "L":
			total += 50
		case "C": 
			if index+1 < len(s) && (string(s[index+1]) == "D" || string(s[index+1]) == "M") {
				total -= 100
				continue
			} 
			total += 100
		case "D":
			total += 500
		case "M":
			total += 1000
		}
	}
	return total
}

func main(){
	fmt.Println(romanToInt("MCMXCIV"))
}