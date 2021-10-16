package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Class int    `json:"class"`
}

func main() {
	f, e := os.Open("टूटा.json")
	if e != nil {
		panic(e)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	var people []Person

	for {
		tok, e := dec.Token()
		if tok == nil || e != nil {
			break
		}
		switch tp := tok.(type) {
		case json.Delim:
			str := string(tp)
			if str == "[" || str == "{" {
				for dec.More() {
					p := Person{}
					e := dec.Decode(&p)
					if e == nil {
						people = append(people, p)
					} else {
						break
					}
				}
			}
		}
	}
	fmt.Println(people)
}
