package gigasecond

// import path for the time package from the standard library
import (
	"math"
	"time"
)

const testVersion = 4

func AddGigasecond(t time.Time) time.Time {
	d := time.Duration(math.Pow10(9)) * time.Second
	newTime := t.Add(d)
	return newTime
}