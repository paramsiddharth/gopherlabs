package main

import "fmt"

func main() {
	var name string
	fmt.Print("आपका नाम – ")
	fmt.Scanf("%s", &name)
	fmt.Printf("नमस्कार, %s ।\n", name)
}
