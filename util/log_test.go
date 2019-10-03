package util

import (
	"errors"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestZap2(t *testing.T) {

	LogInfo("log 初始化成功")
	LogInfo("无法获取网址",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))
	LogDebug("debug")
	LogError("error")
	LogWarn("warn")

	LogErr(errors.New("tt"), zap.String("tag", "eee"))
}

func BenchmarkLogFileSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LogInfo("test size")
	}
}
