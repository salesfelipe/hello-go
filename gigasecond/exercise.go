package gigasecond

import (
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	var gigaSec int64 = addNineZeros(1)

	return t.Add(time.Duration(addNineZeros(gigaSec)))
}

func addNineZeros(seconds int64) int64 {
	return seconds * 1e9
}
