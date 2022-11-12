package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Random string of 6 characters
func RandomName() string {
	return RandomString(6)
}

// Random currency code
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD", "CNY", "CHF", "JPY", "NOK", "SEK", "SGD", "AUD", "GBP", "DKK"}
	length := len(currencies)
	return currencies[rand.Intn(length)]
}
