package singleton

import (
	"testing"
)

func TestIndex(t *testing.T) {
	instance1 := GetInstance()

	hello := "hello"
	world := "world"
	instance1[hello] = world

	instance2 := GetInstance()

	if instance2[hello] != world {
		t.Error("Singleton instance results is different")
	}
}
