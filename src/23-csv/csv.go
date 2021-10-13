package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, e := os.Open("data.csv")
	if e != nil {
		panic(e)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Fprintln(os.Stderr, e)
			break
		}
		fmt.Println(strings.Join(record, ", "))
	}
}
