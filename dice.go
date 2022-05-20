package dice

import (
	"crypto/rand"
)

func RollRecursive() int {
	b := make([]byte, 1)
	rand.Read(b)
	x := b[0] & 7
	if x > 5 {
		return Roll()
	}
	return int(x + 1)
}

func Roll() int {
	b := make([]byte, 6)
	rand.Read(b)
	var sum int
	for _, v := range b {
		sum += int(v)
	}
	return 1 + sum % 6
}