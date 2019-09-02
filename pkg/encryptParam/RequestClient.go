package encryptParam

import (
	"fmt"
	"time"
)

type RequestClient struct {
	RequestId          string    `json:"requestId"`
	TimeStamp          time.Time `json:"timeStamp"`
	EncodeKey          string    `json:"encodeKey"`
	Signature          string    `json:"signature"`
	SignType           string    `json:"signType"`
	RequestJson        string    `json:"requestJson"`
	EncodedRequestJson string    `json:"encodedRequestJson"`
}

func (requestParam *RequestClient) generateSign() error {
	if len(requestParam.EncodedRequestJson) >= 1 {
		return fmt.Errorf("encodedRequestJson %s is not empty", requestParam.EncodedRequestJson)
	}
	if len(requestParam.Signature) >= 1 {
		return fmt.Errorf("signature %s is not empty", requestParam.Signature)
	}
	return nil
}
