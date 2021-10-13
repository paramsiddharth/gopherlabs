package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	num := 1 + 2i

	fmt.Println(num + complex(3, 4))
	fmt.Println(cmplx.Conj(num))
	fmt.Println(1 / num)
}
