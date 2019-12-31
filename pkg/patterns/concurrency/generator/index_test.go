package generator

import (
	"testing"
)

/**
https://github.com/tmrts/go-patterns/blob/master/concurrency/generator.md
Generators yields a sequence of values one at a time.
*/
func TestGenerator(t *testing.T) {
	for i := range Count(1, 20) {
		t.Log("index: ", i)
	}
}
