package logger

import "testing"

func TestZapLog(t *testing.T) {
	InitLog("info", "./zap.log")
	Info("info fang")
	Debug("debug fang")
}
