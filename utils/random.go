package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteByte(alphabet[rand.Intn(len(alphabet))])
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomBalance() int64 {
	return RandomInt(0, 1000)
}
func RandomCurrency() string {
	curr := []string{"USD", "EUR", "KZT", "CAD"}
	return curr[rand.Intn(len(curr))]
}
