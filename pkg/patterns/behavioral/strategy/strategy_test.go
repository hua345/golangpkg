package strategy

import "testing"

/**
https://github.com/tmrts/go-patterns/blob/master/behavioral/strategy.md
Strategy behavioral design pattern enables an algorithm's behavior to be selected at runtime.
*/
func TestStrategy(t *testing.T) {
	multi := Operation{Multiplication{}}
	t.Log(multi.Operate(3, 5))
	add := Operation{Addition{}}
	t.Log(add.Operate(3, 5))
}
