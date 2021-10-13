package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var str = "Param Siddharth  is my name!"
	fmt.Println(str)

	words := strings.Split(str, " ")
	showArray(words)

	words = strings.Fields(str)
	showArray(words)

	words = strings.FieldsFunc(str, func(r rune) bool {
		switch r {
		case ' ', '\n', '\t':
			return true
		}
		return false
	})
	showArray(words)

	words = regexp.MustCompile("\\s").Split(str, -1)
	showArray(words)

	words = regexp.MustCompile("\\s+").Split(str, -1)
	showArray(words)
}

func showArray(words []string) {
	fmt.Println("[ \"" + strings.Join(words, "\", \"") + "\" ]")
}
