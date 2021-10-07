package main

import "fmt"

func main() {
	x := "Param Siddharth"
	msg := `
I am Param!
I am from Bihar, India.`
	fmt.Printf("--------%s--------", x)
	fmt.Println(msg)

	bytesConverted := []byte(x)
	fmt.Printf("%v\n", bytesConverted)
}
