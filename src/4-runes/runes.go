package main

import "fmt"

func main() {
	var x rune
	x = 'प'
	fmt.Printf("%[1]c (%[1]T)\n", x, x)
}
