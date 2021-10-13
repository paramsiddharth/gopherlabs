package main

import "fmt"

func main() {
	const name = "Param"

	fmt.Printf("'% 20s'\n", name)
	fmt.Printf("'% 10s% -10s'\n", name[:len(name)/2], name[len(name)/2:])
	fmt.Printf("'% -20s'\n", name)
}
