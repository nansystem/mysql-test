package common

import (
	"fmt"
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
	return rand.Int63n(max.Unix()-min.Unix()+1) + min.Unix()
}

func RandTime(min, max time.Time) time.Time {
	return time.Unix(RandUnixTime(min, max), 0)
}

func RandTimeSeries(base time.Time, addMin, addMax time.Duration) func() time.Time {
	b := base
	min := int64(addMin)
	max := int64(addMax)
	return func() time.Time {
		b = b.Add(time.Duration(RandNum(min, max)))
		return b
	}
}

var telPres = [...]string{"090", "080", "070", "050"}

// WARNING 電話番号が実在する可能性あり
func RandPhoneNumber() string {
	return fmt.Sprintf("%s-%04d-%04d", telPres[rand.Intn(len(telPres))], RandNum(0, 9999), RandNum(0, 9999))
}

type Period struct {
	Start time.Time
	End   time.Time
}

func RandPeriod(baseMin time.Time, baseMax time.Time, addMinDay, addMaxDay int64) Period {
	addDay := RandNum(addMinDay, addMaxDay)
	start := RandTime(baseMin, baseMax)
	end := start.AddDate(0, 0, int(addDay))
	return Period{start, end}
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(min, max int64) string {
	b := make([]byte, RandNum(min, max))

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

type ValWeight struct {
	Val    interface{}
	Weight int
}

func RandValWeight(vws []ValWeight) interface{} {
	totalWeight := 0
	for _, vw := range vws {
		totalWeight += vw.Weight
	}
	rnd := rand.Int() % totalWeight
	for _, vw := range vws {
		if rnd < vw.Weight {
			return vw.Val
		}
		rnd -= vw.Weight
	}
	return vws[len(vws)-1]
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
