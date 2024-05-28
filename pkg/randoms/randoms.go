package randoms

import (
	"math/rand"
	"time"
)

const (
	R = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890abcdefghijklmnopqrstuvwxyz"
)

func Alphanumeric(n uint) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var s string
	for i := uint(0); i < n; i++ {
		v := string(R[r.Intn(len(R)-1)])
		s += v
	}
	return s
}
