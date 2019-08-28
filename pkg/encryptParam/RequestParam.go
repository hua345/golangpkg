package encryptParam

import "time"

type RequestParam struct {
	RequestId          string    `json:"requestId"`
	TimeStamp          time.Time `json:"timeStamp"`
	EncodeKey          string    `json:"encodeKey"`
	Signature          string    `json:"signature"`
	RequestJson        string    `json:"requestJson"`
	EncodedRequestJson string    `json:"encodedRequestJson"`
}
