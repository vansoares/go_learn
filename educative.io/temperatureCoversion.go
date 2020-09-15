package main

import "fmt"

type Celsius float32
type Fahrenheit float32

func toFahrenheit(t Celsius) Fahrenheit {
	var temperature Fahrenheit
	temperature = Fahrenheit((t * 9 / 5) + 32)
	return temperature
}

func main() {
	var c Celsius
	c = 100

	f := toFahrenheit(c)
	fmt.Println(f)
}
