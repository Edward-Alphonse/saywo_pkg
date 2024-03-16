package logs

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/Edward-Alphonse/saywo_pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const logPath = "/tmp/log/server/saywo"

var logger *zap.Logger

func Init(config *Config) {
	if utils.IsDevEnv() {
		initDevLogger()
		return
	}
	initProductionLogger(config)
}

func getEncoder() zapcore.Encoder {
	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		_, file, line, ok := runtime.Caller(6)
		if ok {
			path := fmt.Sprintf("%s:%d", file, line)
			path = trimmedPath(path)
			enc.AppendString("[" + path + "]")
		} else {
			enc.AppendString("[" + caller.TrimmedPath() + "]")
		}

	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.CallerKey = "caller_line"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = customCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func initDevLogger() {
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, os.Stdout, zapcore.InfoLevel)
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.DPanicLevel))
}

func initProductionLogger(config *Config) {
	writeSyncer := getDefaultLogWriter()
	if config != nil {
		writeSyncer = getCustomLogWriter(config)
	}
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.DPanicLevel))
}

func getDefaultLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logPath + getLogFileName(),
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getCustomLogWriter(config *Config) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.Path + getLogFileName(),
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getLogFileName() string {
	return fmt.Sprintf("%s.log", time.Now().Format("20060102150405"))
}

func Debug(msg string, fields map[string]any) {
	zapFields := getZapFileds(fields)
	logger.Debug(msg, zapFields...)
}

func Info(msg string, fields map[string]any) {
	zapFields := getZapFileds(fields)
	logger.Info(msg, zapFields...)
}

func Error(msg string, fields map[string]any) {
	zapFields := getZapFileds(fields)
	logger.Error(msg, zapFields...)
}

func Warn(msg string, fields map[string]any) {
	zapFields := getZapFileds(fields)
	logger.Warn(msg, zapFields...)
}

func Panic(msg string, fields map[string]any) {
	zapFields := getZapFileds(fields)
	logger.Panic(msg, zapFields...)
}

func DPanic(msg string, fields map[string]any) {
	zapFields := getZapFileds(fields)
	logger.DPanic(msg, zapFields...)
}

func Fatal(msg string, fields map[string]any) {
	zapFields := getZapFileds(fields)
	logger.Fatal(msg, zapFields...)
}

func getZapFileds(fields map[string]any) []zap.Field {
	list := make([]zap.Field, 0)
	for key, value := range fields {
		list = append(list, zap.Any(key, value))
	}
	return list
}

func trimmedPath(path string) string {
	idx := len(path)
	for i := 0; i < 3; i++ {
		idx = strings.LastIndexByte(path[:idx], '/')
		if idx == -1 {
			return path
		}
	}
	file := path[idx+1:]
	return file
}

/* 适配旧项目*/

// InfoByArgs 通过参数输出日志
func InfoByArgs(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	logger.Info(s)
}

// ErrorByArgs 通过参数输出错误日志
func ErrorByArgs(format string, args ...interface{}) {
	logger.Error(fmt.Sprintf(format, args...))
}

// WarnByArgs 通过参数输出警告日志
func WarnByArgs(format string, args ...interface{}) {
	logger.Warn(fmt.Sprintf(format, args...))
}

// DebugByArgs 通过参数输出debug日志
func DebugByArgs(format string, args ...interface{}) {
	logger.Debug(fmt.Sprintf(format, args))
}

//func ErrorKv(key string, value string, format string, args ...interface{}) {
//	logger.Error(fmt.Sprintf(format, args), zap.String(key, value))
//}

func ErrorKv(key string, value string, format string, args ...interface{}) {
	logger.Error(fmt.Sprintf(format, args), zap.String(key, value))

}

func InfoKv(key string, value string, format string, args ...interface{}) {
	logger.Info(fmt.Sprintf(format, args), zap.String(key, value))
}

func WarnKv(key string, value string, format string, args ...interface{}) {
	logger.Warn(fmt.Sprintf(format, args), zap.String(key, value))
}
