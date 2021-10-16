package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// अस्थायी नत्थी
	f, e := ioutil.TempFile("", "अस्थायी")
	if e != nil {
		panic(e)
	}
	defer os.Remove(f.Name())

	fmt.Println("File: ", f.Name())

	// अस्थायी संचय
	fo, e := ioutil.TempDir("", "अस्थायी")
	if e != nil {
		panic(e)
	}
	defer os.Remove(fo)

	fmt.Println("Folder: ", fo+string(filepath.Separator))
}
