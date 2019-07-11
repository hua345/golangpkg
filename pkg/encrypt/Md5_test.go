package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"testing"
)

func TestMd5(t *testing.T) {
	//原始数据
	var data = []byte("fang")
	dataMd5 := "0b527ad35a7983fa5c9abdf31825c3cb"
	//md5
	hash := md5.New()
	hash.Write(data)
	encryptedData := hash.Sum(nil)
	t.Log("md5:", hex.EncodeToString(encryptedData))
	if hex.EncodeToString(encryptedData) != dataMd5 {
		t.Error("Md5 结果不一致")
	}
}
