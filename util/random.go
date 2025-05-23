package util

import (
	"time"

	"math/rand"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyz")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func Randomusername() string {
	return RandomString(6)
}
