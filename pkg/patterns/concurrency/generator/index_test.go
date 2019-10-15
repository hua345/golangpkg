package generator

import (
	"fmt"
	"testing"
)

func TestGenerator(t *testing.T) {
	for i := range Count(1, 99) {
		fmt.Println("index: ", i)
	}
}
