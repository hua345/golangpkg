package util

import (
	"github.com/satori/go.uuid"
	"strings"
)

func GetUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}
func GetUUID32() string {
	uuidStr := GetUUID()
	uuidStr = strings.Replace(uuidStr, "-", "", -1)
	return uuidStr
}
