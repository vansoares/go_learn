package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("BOOL!")
		case int:
			fmt.Println("INT!!")
		default:
			fmt.Printf("WHAT?!")
		}
	}

	whatAmI(true)
	whatAmI(5)
}
