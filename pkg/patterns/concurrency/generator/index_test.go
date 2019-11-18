package generator

import (
	"testing"
)

func TestGenerator(t *testing.T) {
	for i := range Count(1, 99) {
		t.Log("index: ", i)
	}
}
