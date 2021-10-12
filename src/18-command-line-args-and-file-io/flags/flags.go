package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "", "Your name")
	flag.Parse()

	if name == "" {
		fmt.Println("Hello, World!")
	} else {
		fmt.Println("Hello, " + name + "!")
	}
}
