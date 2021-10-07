package main

import (
	"fmt"
	"strings"
)

func main() {
	var txt string = "मुझे समोसे पसंद हैं ।"
	fmt.Printf("पाठ्यांश – %s\n", txt)
	fmt.Printf("क्या यहाँ पर \"समोसे\" हैं ? %s ।\n", func() string {
		if strings.ContainsAny(txt, "समोसे") {
			return "हाँ"
		}
		return "नहीं"
	}())
}
