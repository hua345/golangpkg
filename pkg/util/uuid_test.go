package util

import (
	"testing"
)

// 单元测试
// go test -v
func TestGetUUID(t *testing.T) {
	t.Log("UUIDv4: " + GetUUID())
	uuidStr := GetUUID()
	t.Log("UUIDv4: " + uuidStr)
	if len(uuidStr) != 36 {
		t.Error(`UUID {uuidStr} length != 36`)
	}
}
func TestUUID32(t *testing.T) {
	t.Log("UUID32: ", GetUUID32())
}
