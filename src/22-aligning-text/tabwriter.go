package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

var students []map[string]string = []map[string]string{
	{
		"name":  "Param",
		"class": "15",
		"id":    "1",
	},
	{
		"name":  "Kabir",
		"class": "9",
		"id":    "2",
	},
	{
		"name":  "Hanuman",
		"class": "200",
		"id":    "3",
	},
}

func main() {
	writer := tabwriter.NewWriter(os.Stdout, 10, 4, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(writer, "Name\tClass\tRoll Number\t")
	for _, v := range students {
		fmt.Fprintf(writer, "%s\t%s\t%s\t\n", v["name"], v["class"], v["id"])
	}
	writer.Flush()
}
