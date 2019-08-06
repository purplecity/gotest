package main

import (
	"crypto/rand"
	"fmt"
	r "math/rand"
	"time"
)



// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes(n int) string {
	alphabets := []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)
	var bytes = make([]byte, n)
	var randBy bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randBy = true
	}
	for i, b := range bytes {
		if randBy {
			bytes[i] = alphabets[r.Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return string(bytes)
}

func main() {
	fmt.Println(string(RandomCreateBytes(6)))
}