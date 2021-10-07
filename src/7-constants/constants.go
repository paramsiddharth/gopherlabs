package main

import "fmt"

const x = 1
const y float64 = 1.0
const z = "Hi"

const (
	a        = 1
	b        = '2'
	c string = "Three"
)

func main() {
	fmt.Println(x, y, z, a, b, c)
}
