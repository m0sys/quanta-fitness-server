package random

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// String generate a random string of size `n`.
func String(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)

	}
	return sb.String()
}

// Float generate a random float64 between in [min, max].
func Float(min, max float64) float64 {
	return min + rand.Float64()*(max-min+1)
}

// Int generate a random int between in [min, max].
func Int(min, max int64) int {
	return int(min + rand.Int63n(max-min+1))

}

// Weight generate a random weight in kg.
func Weight() float64 {
	return roundToTwoDecimalPlaces(Float(10, 198.0))
}

// RestTime generate a random rest time in seconds.
func RestTime() float64 {
	return roundToTwoDecimalPlaces(Float(30.0, 150.0))
}

// RepCount generate a random rep count.
func RepCount() int {
	return Int(5, 15)
}

// NumSets generate a random number of sets.
func NumSets() int {
	return Int(1, 25)
}

// Height generate a random height in m.
func Height() float64 {
	return Float(1.0, 2.0)
}

// Email generate a random email.
func Email() string {
	return fmt.Sprintf("%s@email.com", String(10))
}

func roundToTwoDecimalPlaces(num float64) float64 {
	return math.Round(num*100) / 100
}
