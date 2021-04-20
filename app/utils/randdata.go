package utils

import (
	"math/rand"
	"time"
)

func RandNum(min, max int64) int64 {
	return rand.Int63n(max-min+1) + min
}

func RandUnixTime(min, max time.Time) int64 {
	return rand.Int63n(max.Unix()-min.Unix()+1) + max.Unix()
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
