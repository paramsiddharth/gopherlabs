package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sumofsquares(1, 2, 3))
}

func sumofsquares(nums ...int32) int32 {
	sum := int32(0)
	for i := 0; i < len(nums); i++ {
		sum += int32(math.Pow(float64(nums[i]), 2))
	}
	return sum
}
