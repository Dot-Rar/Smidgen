package utils

import "math/rand"

var(
	charset = "abcdefghijklmnopqrstuvwxyz1234567890"
)

func GenerateId(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
