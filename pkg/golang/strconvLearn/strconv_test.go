package strconvLearn

import (
	"strconv"
	"testing"
)

func TestParseInt(t *testing.T) {
	//Bit sizes 0, 8, 16, 32, and 64
	// correspond to int, int8, int16, int32, and int64.
	int64Result, err := strconv.ParseInt("314520", 10, 64)
	if err != nil {
		t.Log(err)
	}
	t.Log(int64Result)
}
func TestAtoi(t *testing.T) {
	intResult, err := strconv.Atoi("314520")
	if err != nil {
		t.Log(err)
	}
	t.Log(intResult)
}
func TestItoa(t *testing.T) {
	str := strconv.Itoa(618520)
	t.Log(str)
}
func TestFormatInt(t *testing.T) {
	str := strconv.FormatInt(618520, 10)
	t.Log(str)
}
