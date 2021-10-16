package main

import (
	"fmt"
	"os"
)

func main() {
	f, e := os.Open("नमस्कार.txt")
	if e != nil {
		panic(e)
	}

	fi, e := f.Stat()
	if e != nil {
		panic(e)
	}

	fmt.Printf("Name: %s\n", fi.Name())
	fmt.Printf("Size: %d\n", fi.Size())
	fmt.Printf("Directory: %t\n", fi.IsDir())
	fmt.Printf("Mode: %s\n", fi.Mode())
}
