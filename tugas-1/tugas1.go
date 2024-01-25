package main

import (
	"fmt"
	"math"
)

func division(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("Unable to do divide by zero")
	}
	return a / b, nil
}

func logarithm(argument float64) (float64, error) {
	if argument <= 0 {
		return 0, fmt.Errorf("Unable to do log with %f", argument)
	}
	return math.Log(float64(argument)), nil
}

func main() {
	// fmt.Println(quote.Go())
	fmt.Println("Let's Go", "Print")
	formattedString := fmt.Sprintln("Let us Go", "Sprint")
	fmt.Print((formattedString))

	c, _ := division(2, 3)
	d, er := division(4, 0)
	fmt.Println(c, d, er)

	var f, g float64
	// var _ error
	f, _ = logarithm(100)
	g, er = logarithm(-10)
	fmt.Println(f, g, er)
}
