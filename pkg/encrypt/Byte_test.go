package encrypt

import (
	"github.com/imroc/biu"
	"testing"
)

/**
打印二进制01格式
*/
func TestA(t *testing.T) {
	// 1个字节=8个二进制位,每种数据类型占用的字节数都不一样
	var aa uint8 = 179
	// 10110011
	t.Log(biu.ToBinaryString(aa))
	// G(X)=X4+X3+1 CRC二进制串11001
	var bb uint8 = 25
	// 00011001
	t.Log(biu.ToBinaryString(bb))
	/**
	  将某一位设置为1，例如设置第8位，从右向左数需要偏移7位,注意不要越界
	  1<<7=1000 0000 然后与a逻辑或|,偏移后的第8位为1，逻辑|运算时候只要1个为真就为真达到置1目的
	*/
	cc := aa | (1 << 7)
	// 11111000
	t.Log(biu.ToBinaryString(cc))
	/**
	  将某一位设置为0，例如设置第4位，从右向左数需要偏移3位,注意不要越界
	  1<<3=0000 1000 异或 01111000 = 01110000
	*/
	dd := aa ^ (1 << 3)
	//
	t.Log(biu.ToBinaryString(dd))
}

/**
异或运算
*/
func TestXOR(t *testing.T) {
	t.Log(0 ^ 0)
	t.Log(0 ^ 1)
	t.Log(1 ^ 0)
	t.Log(1 ^ 1)
	if 0^0 != 0 {
		t.Error("0^0 != 0")
	}
	if 0^1 != 1 {
		t.Error("0^1 != 1")
	}
	if 1^0 != 1 {
		t.Error("1^0 != 1")
	}
	if 1^1 != 0 {
		t.Error("1^1 != 0")
	}
}
