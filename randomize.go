package randomize

import (
  "time"
  "math/rand"
)
// Random is a func to randomize number with custom range
func Random(randomRange int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(randomRange)
}