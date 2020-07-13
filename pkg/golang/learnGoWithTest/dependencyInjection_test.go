package learnGoWithTest

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

/**
我们的函数不需要关心在哪里打印，以及如何打印，所以我们应该接收一个接口，而非一个具体的类型。
fmt.Fprintf 允许传入一个 io.Writer 接口，我们知道 os.Stdout 和 bytes.Buffer 都实现了它
*/
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func TestGreet(t *testing.T) {
	Greet(os.Stdout, "Fang")
	buffer := bytes.Buffer{}
	Greet(&buffer, "Fang")

	got := buffer.String()
	want := "Hello, Fang"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
