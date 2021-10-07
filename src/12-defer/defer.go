package main

import "fmt"

func showNums() {
	defer fmt.Println()

	for i := 9; i > 0; i-- {
		defer fmt.Printf("%d ", i)
	}
}

func main() {
	defer showNums()
	defer fmt.Println(", World!")

	fmt.Print("Hello")
}
