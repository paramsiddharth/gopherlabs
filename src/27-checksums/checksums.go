package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, e := ioutil.ReadFile("data.txt")
	if e != nil {
		panic(e)
	}

	contents := string(bytes)

	fmt.Println("\"" + contents + "\"")

	md5sum := md5.Sum([]byte(contents))
	sha256sum := sha256.Sum256(bytes)

	fmt.Printf("MD5: %x\n", md5sum)
	fmt.Printf("SHA-256: %x\n", sha256sum)
}
