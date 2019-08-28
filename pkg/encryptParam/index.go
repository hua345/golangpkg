package encryptParam

import (
	"encoding/base64"
	"encoding/json"
	"golangpkg/pkg/encrypt"
	"golangpkg/pkg/util"
	"time"
)

func Encrypy(param map[string]interface{}) (*RequestParam, error) {
	requestParam := &RequestParam{}
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
	data, err := json.Marshal(param)
	if err != nil {
		return requestParam, err
	}
	requestParam.RequestJson = string(data)
	// 获取加密后的参数
	data, err = json.Marshal(param)
	if err != nil {
		return requestParam, err
	}
	xpass, err := encrypt.AesEncrypt(data, []byte(data))
	if err != nil {
		return requestParam, err
	}

	pass64Str := base64.StdEncoding.EncodeToString(xpass)
	requestParam.EncodedRequestJson = pass64Str
	return requestParam, nil
}
