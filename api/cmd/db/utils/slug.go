package utils

import (
	"crypto/rand"
	"github.com/gosimple/slug"
	"math/big"
)

func generateRandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		r, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[r.Int64()]
	}
	return string(b)
}

func Slugify(s string) string {
	return slug.Make(s + "-" + generateRandomString(8))
}
