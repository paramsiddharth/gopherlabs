package main

import "fmt"

func main() {
	i := 1

	// While loop construct
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println()
	i = 0

	// Conditionless loop (forever)
	for {
		fmt.Println(i)
		i++
		if i >= 10 {
			break
		}
	}

	fmt.Println()

	// For loop (syntactic sugar)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
