package encrypt

import (
	"fmt"
	"testing"
)

func TestPKCS5Padding(t *testing.T) {
	data := []byte("fangfang")
	fmt.Printf("%v\n", data)
	data = PKCS5Padding(data, 16)
	fmt.Printf("%v\n", data)
}

func TestZeroPadding(t *testing.T) {
	data := []byte("fangfang")
	fmt.Printf("%v\n", data)
	data = ZeroPadding(data, 16)
	fmt.Printf("%v\n", data)
}
