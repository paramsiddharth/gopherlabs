package randnum

import (
	"fmt"
	"math/rand"
	"time"
)

func MathRand() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
}
