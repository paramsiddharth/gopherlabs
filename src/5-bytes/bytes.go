package main

import "fmt"

func main() {
	var x byte
	x = byte('Q')
	fmt.Printf("%c (%[1]T)\n", x)
}
