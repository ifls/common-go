package utils

import (
	"go.uber.org/zap/zapcore"
	"testing"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05.000000"
)

func Judge(result bool, t *testing.T) {
	if result {
		t.Log("result is true or right")
	} else {
		t.Fatal("result is false or error")
	}
}

func RunDelay(t time.Timer, delay int, fun func()) {
	t.Stop()
	t.Reset(time.Duration(delay) * time.Millisecond)
	<-t.C
	fun()
}

func GetTime() string {
	return time.Now().Format("2006-01-02 15:04:05.000000")
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000000"))
}
