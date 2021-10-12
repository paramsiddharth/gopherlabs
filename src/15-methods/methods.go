package main

import "github.com/paramsiddharth/gopherlabs/src/15-methods/student"

func main() {
	param := student.Student{
		ID:     1,
		Name:   "Param Siddharth",
		Gender: student.Male,
	}

	param.Display()
}
