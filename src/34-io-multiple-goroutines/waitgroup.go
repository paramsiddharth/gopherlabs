package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	f, e := os.Create("data")
	if e != nil {
		panic(e)
	}
	defer f.Close()

	data := []string{
		"Param",
		"Kabir",
		"Hanuman",
	}

	for _, name := range data {
		wg.Add(1)
		go func(t string) {
			fmt.Fprintln(f, t)
			fmt.Println(t)
			wg.Done()
		}(name)
	}

	wg.Wait()
}
