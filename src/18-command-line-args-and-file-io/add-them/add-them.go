package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Invalid arguments!")
		os.Exit(1)
		// panic(e)
	}

	var nums [2]int
	for i, num := range os.Args {
		if i == 0 {
			continue
		}
		n, e := strconv.Atoi(num)
		if e != nil {
			fmt.Fprintf(os.Stderr, "Invalid number '%s'!\n", num)
			os.Exit(1)
			// panic(e)
		}
		nums[i-1] = n
	}

	sum := nums[0] + nums[1]
	fmt.Println(sum)

	cwd, _ := os.Getwd()
	filename := path.Join(cwd, "output")
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open 'output' file for writing!\n")
		os.Exit(1)
		// panic(err)
	}
	defer file.Close()

	file.WriteString(fmt.Sprintln(sum))
}
