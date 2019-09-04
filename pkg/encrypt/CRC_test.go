package encrypt

import (
	"bytes"
	"encoding/binary"
	"github.com/imroc/biu"
	"golangpkg/pkg/util"
	"testing"
)

// 生成多项式:G(X)=X4+X3+1，要求出二进制序列10110011的CRC校验码。
//（1）G(X)=X4+X3+1,二进制比特串为11001;(有X的几次方，对应的2的几次方的位就是1)
//（2）因为校验码4位，所以10110011后面再加4个0，得到101100110000，用“模2除法”(其实就是亦或^)即可得出结果；
// （3）CRC^101100110000得到101100110100。发送到接收端；
// （4）接收端收到101100110100后除以11001(以“模2除法”方式去除),余数为0则无差错；
// 101100110000 XOR 00011001
// 101100110000 XOR
// 11001
// 011110        XOR
//  11001
//  0011111      XOR
//    11001
//    0011000    XOR
//      11001
//      0000100  XOR
// 获取到4为CRC校验码0100
func TestCRC(t *testing.T) {
	// 1个字节=8个二进制位,每种数据类型占用的字节数都不一样
	var aa uint8 = 179
	// 10110011
	t.Log(biu.ToBinaryString(aa))
	// G(X)=X4+X3+1 CRC二进制串11001
	var bb uint8 = 1<<4 + 1<<3 + 1
	// 00011001
	t.Log(biu.ToBinaryString(bb))
}

// redis插槽
// 192.168.137.128:6380> cluster keyslot name
//(integer) 5798
func TestCRC16(t *testing.T) {
	data := []byte("name")
	checksum := CRC16CheckSum(data)
	t.Log(checksum % 16384)
	t.Logf("checksum: %X", checksum)
	t.Log("Little-Endian", biu.ToBinaryString(checksum))
	int16buf := new(bytes.Buffer)
	//  Little-Endian就是低位字节排放在内存的低地址端，高位字节排放在内存的高地址端。
	//  Big-Endian就是高位字节排放在内存的低地址端，低位字节排放在内存的高地址端。
	binary.Write(int16buf, binary.LittleEndian, checksum)
	t.Log("Little-Endian", biu.ToBinaryString(int16buf.Bytes()))
	t.Logf("Little-Endian checksum: %+X", int16buf.Bytes())

	t.Logf("origin data: %X", data)
	data = append(data, int16buf.Bytes()...)

	t.Logf("output-after: %X", data)
}
func TestCRC32(t *testing.T) {
	t.Log(CRC32Hash(util.GetUUID32()) % 32)
	for i := 0; i < 100; i++ {
		t.Log(CRC32Hash(util.GetUUID32()) % 10)
	}
}
