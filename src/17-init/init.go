package main

import (
	"fmt"
	"time"
)

var weekday string

func main() {
	fmt.Println("Today is " + weekday + "!")
}

func init() {
	weekday = time.Now().Weekday().String()
}
