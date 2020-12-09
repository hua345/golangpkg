package encryptParam

import (
	"encoding/base64"
	"encoding/json"
	"github.com/hua345/golangpkg/pkg/encrypt"
	"github.com/hua345/golangpkg/pkg/util"
	"time"
)

func Encrypt(param map[string]interface{}) (*RequestClient, error) {
	requestParam := &RequestClient{}
	requestParam.RequestId = util.GetUUID32()
	requestParam.TimeStamp = time.Now()
	// 需要脱敏的字段
	// 随机生成AES密钥
	aesKey := util.GetUUID32()
	// 对加密密钥进行rsa加密
	encodeKey, err := encrypt.RsaEncrypt([]byte(aesKey), encrypt.PublicKey)
	if err != nil {
		return requestParam, err
	}
	requestParam.EncodeKey = encodeKey
	// 参数
	paramJsonData, err := json.Marshal(param)
	if err != nil {
		return requestParam, err
	}
	requestParam.RequestJson = string(paramJsonData)

	// 获取加密后的参数
	xpass, err := encrypt.AesEncryptCBC(paramJsonData, []byte(aesKey))
	if err != nil {
		return requestParam, err
	}

	pass64Str := base64.StdEncoding.EncodeToString(xpass)
	requestParam.EncodedRequestJson = pass64Str
	requestParam.RequestJson = ""
	return requestParam, nil
}
