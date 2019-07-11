package jwt

import (
	"testing"
)

// 单元测试
// go test -v
func TestCreateMapToken(t *testing.T) {
	helloMap := map[string]string{
		"name":   "fang",
		"userId": "001",
	}
	tokenStr := CreateMapToken(helloMap)
	t.Log("Token: " + tokenStr)
	tokenStrWithKey := CreateMapToken(helloMap)
	t.Log("TokenStrWithKey: " + tokenStrWithKey)
}

func TestParseMapToken(t *testing.T) {
	InitJwt("fangfang")
	helloMap := map[string]string{
		"name":   "fang",
		"userId": "001",
	}
	tokenStr := CreateMapToken(helloMap)
	t.Log("Token: " + tokenStr)
	resultMap, ok := ParseMapToken(tokenStr)
	if ok {
		for key, value := range resultMap {
			t.Log("key:" + key + " value: " + value)
		}
	}
}
