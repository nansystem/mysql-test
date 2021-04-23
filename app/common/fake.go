package common

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandNum(min, max int64) int64 {
	return rand.Int63n(max-min+1) + min
}

func RandUnixTime(min, max time.Time) int64 {
	return rand.Int63n(max.Unix()-min.Unix()+1) + max.Unix()
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(min, max int64) string {
	b := make([]byte, RandNum(min, max))

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

type IndexChunk struct {
	From, To int
}

// IndexChunks for Error 1390: Prepared statement contains too many placeholders
func IndexChunks(length int, chunkSize int) <-chan IndexChunk {
	ch := make(chan IndexChunk)
	go func() {
		defer close(ch)
		for i := 0; i < length; i += chunkSize {
			idx := IndexChunk{i, i + chunkSize}
			if length < idx.To {
				idx.To = length
			}
			ch <- idx
		}
	}()
	return ch
}
