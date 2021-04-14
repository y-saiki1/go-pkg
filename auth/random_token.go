package auth

import (
	"crypto/rand"
	"encoding/binary"
	"strconv"
)

func CreateRandomString() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}
