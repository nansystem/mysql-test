package common

import "time"

var (
	jp *time.Location
)

func init() {
	jp = time.FixedZone("Asia/Tokyo", 9*60*60)
}

func JP() *time.Location {
	return jp
}
