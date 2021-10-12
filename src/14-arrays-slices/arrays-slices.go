package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [...]int{
		5, 6, 7, 8, 9,
		10,
		11, 12, 13,
	}

	show(arr1)
	show(arr2)

	slice1 := []int{0, 0, 0}
	var slice2 []int = make([]int, 3)
	copy(slice2, arr2[:3])

	show(slice1)
	show(slice2)
}

func show(x interface{}) {
	fmt.Printf("%[1]T %[1]v\n", x)
}
