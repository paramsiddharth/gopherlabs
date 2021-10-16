package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	ID      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Class   string   `xml:"class"`
}

func main() {
	f, e := os.Open("data.xml")
	if e != nil {
		panic(e)
	}
	defer f.Close()

	decoder := xml.NewDecoder(f)
	var people []Person

	for {
		tok, _ := decoder.Token()
		// fmt.Printf("%[1]T:\t%[1]v\n", tok)
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "person" {
				var p Person
				decoder.DecodeElement(&p, &tp)
				people = append(people, p)
			}
		}
	}
	fmt.Println(people)
}
