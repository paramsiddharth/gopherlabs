package randnum

import (
	"crypto/rand"
	"encoding/binary"
)

func CryptoRand() {
	var n int32
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	println(n)
}
