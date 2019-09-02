package encryptParam

import (
	"encoding/json"
	"testing"
)

func TestEncrypy(t *testing.T) {
	param := map[string]interface{}{
		"name": "fang",
		"age":  20,
	}
	requestParam, err := Encrypt(param)
	if err != nil {
		t.Error(err)
	}
	result, err := json.MarshalIndent(requestParam, "", "    ")
	t.Log(string(result))
}
