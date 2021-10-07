package main

import "fmt"

func main() {
	var x int32
	var y float64
	var z rune
	var a byte
	var b string
	var c bool
	var d complex128

	fmt.Printf("%v (%[1]T)\n", x)
	fmt.Printf("%v (%[1]T)\n", y)
	fmt.Printf("%v (%[1]T)\n", z)
	fmt.Printf("%v (%[1]T)\n", a)
	fmt.Printf("%v (%[1]T)\n", b)
	fmt.Printf("%v (%[1]T)\n", c)
	fmt.Printf("%v (%[1]T)\n", d)
}
