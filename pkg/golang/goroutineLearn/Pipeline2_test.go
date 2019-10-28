package goroutineLearn

import "testing"

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int, t *testing.T) {
	for v := range in {
		t.Log(v)
	}
}

func TestPipeline2(t *testing.T) {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares, t)
}
