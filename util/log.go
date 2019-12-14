package util

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"net"
	"os"
	"runtime"
	"runtime/debug"
)

var logger *zap.Logger

const (
	LogTagReason = "reason"
)

var logDir string

func init() {
	switch runtime.GOOS {
	case "darwin":
		logDir = "/Users/ifls/Downloads/logs/"
	default:
		logDir = "/data/logs"
	}
	Initdefault()
	//InitLogTest(nil, nil)
}

func InitLogTest(lw io.Writer, conn net.Conn) {
	hook := lumberjack.Logger{
		Filename:   logDir + "dev_test.log", // 日志文件路径
		MaxSize:    1,                       // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                      // 日志文件最多保存多少个备份
		MaxAge:     7,                       // 文件最多保存多少天
		//Compress:   true,                     // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 小写编码器
		EncodeTime:     TimeEncoder,                      // yyyy-mm-dd hh-mm-ss.xxxxxx
		EncodeDuration: zapcore.SecondsDurationEncoder,   //
		EncodeCaller:   nil,                              // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook), zapcore.AddSync(lw)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serverType", "logTest"), zap.String("server addr", conn.LocalAddr().String()))
	// 构造日志
	l := zap.New(core, caller, development, filed)

	if l != nil {
		InitLogger(l)
	} else {
		log.Fatal("init logger fail")
	}
}

func Initdefault() {
	hook := lumberjack.Logger{
		Filename:   logDir + "dev_test.log", // 日志文件路径
		MaxSize:    1,                       // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                      // 日志文件最多保存多少个备份
		MaxAge:     7,                       // 文件最多保存多少天
		//Compress:   true,                     // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 小写编码器
		EncodeTime:     TimeEncoder,                      // yyyy-mm-dd hh-mm-ss.xxxxxx
		EncodeDuration: zapcore.SecondsDurationEncoder,   //
		EncodeCaller:   zapcore.ShortCallerEncoder,       // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),                                        // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	callerSkip := zap.AddCallerSkip(1)
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	//filed := zap.Fields(zap.String("serverType", "logdefault"))
	// 构造日志
	l := zap.New(core, caller, callerSkip, development)

	if l != nil {
		InitLogger(l)
	} else {
		log.Fatal("init logger fail")
	}
}

/***********************log API ***********************/

func InitLogger(l *zap.Logger) bool {
	if l != nil {
		if logger != nil {
			logger.Info("default logger changed successful")
		} else {
			l.Info("default logger inited successful")
		}
		logger = l
		return true
	} else {
		if logger != nil {
			logger.Info("params l is nil")
		} else {
			log.Printf("default logger is nil && params logger is nil\n")
		}
		return false
	}
}

func LogDebug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
	if err := logger.Sync(); err != nil {
		log.Println(err)
	}

}

func LogInfo(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
	if err := logger.Sync(); err != nil {
		log.Println(err)
	}
}

func LogWarn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
	if err := logger.Sync(); err != nil {
		log.Println(err)
	}
}

func LogError(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
	if err := logger.Sync(); err != nil {
		log.Println(err)
	}
}

func LogFatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Println(err)
		}
	}()
}

func LogDebugf(msg string, fields ...interface{}) {
	logger.Sugar().Debugf(msg, fields...)
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Println(err)
		}
	}()
}

func LogErr(err error, fields ...zap.Field) {
	if err != nil {
		LogStack()
		logger.Error("error = "+err.Error(), fields...)
		if err := logger.Sync(); err != nil {
			log.Println(err)
		}
	}
}

func Log(level int, msg string, fields ...zap.Field) {
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Println(err)
		}
	}()

	switch zapcore.Level(level) {
	case zapcore.DebugLevel:
		logger.Debug(msg, fields...)
	case zapcore.InfoLevel:
		logger.Info(msg, fields...)
	case zapcore.WarnLevel:
		logger.Warn(msg, fields...)
	case zapcore.ErrorLevel:
		logger.Error(msg, fields...)
	case zapcore.PanicLevel:
		logger.Panic(msg, fields...)
	case zapcore.FatalLevel:
		logger.Fatal(msg, fields...)
	}
}

func LogErrAndExit(err error, fields ...zap.Field) {
	if err != nil {
		defer func() {
			if err := logger.Sync(); err != nil {
				log.Println(err)
			}
		}()
		logger.Fatal("FATAL="+err.Error(), fields...)
	}
}

func DevInfo(format string, v ...interface{}) {
	log.Printf(format, v...)
}

/***********************log API ***********************/
func LogCaller() {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		LogInfo(fmt.Sprintf("file:line=%s:%d func: %s\n", file, line, runtime.FuncForPC(pc).Name()))
	} else {
		LogInfo("log caller error")
	}
}

func LogStack() {
	LogInfo(fmt.Sprintf("traceback:\n%s", debug.Stack()))
}
