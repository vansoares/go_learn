package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewReader(os.Stdin)
	// Declare second integer, double, and String variables.
	var first uint64
	var second float64
	var third string
	// Read and save an integer, double, and String to your variables.
	msg1, _ := scanner.ReadString('\n')
	msg2, _ := scanner.ReadString('\n')
	msg3, _ := scanner.ReadString('\n')
	fmt.Println(msg1, msg2, msg3)
	// Print the sum of both integer variables on a new line.
	first, _ = strconv.ParseUint(strings.Replace(msg1, " ", "", 5), 10, 64)
	second, _ = strconv.ParseFloat(msg2, 64)
	third = msg3

	fmt.Println(i, first)
	fmt.Println(d, second)
	fmt.Println(s, third)
	// Print the sum of the double variables on a new line.

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.

}
