 <= 0 {
		return 0, fmt.Errorf("Unable to do log with %f", argument)
	}
	return math.Log(float64(argument)), nil
}

func main() {
	// fmt.Println(quote.Go())
	fmt.Println("Let's Go", "Print")
	formattedString := fmt.Sprintln("Let us Go", "Sprint")
	fmt.Print((formattedString))

	c, _ := division(2