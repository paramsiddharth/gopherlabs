package main

import "fmt"

type person struct {
	name string
	age  uint16
}

func main() {
	var param person

	// Way 1
	param = person{"Param", 20}

	// Way 2
	param = person{
		name: "Param",
		age:  20,
	}

	// Way 3
	param = person{name: "Param", age: 20}

	var paramPointer *person

	paramPointer = &param
	paramPointer = &person{
		name: "Param",
		age:  20,
	}

	fmt.Println(param)
	fmt.Println(*paramPointer)
}
