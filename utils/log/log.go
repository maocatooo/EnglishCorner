package log

import (
	"EnglishCorner/utils/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type log struct {
	logObj                *zap.SugaredLogger
	Infof, Debugf, Errorf func(string, ...interface{})
	Info, Debug, Error    func(...interface{})
}

var logobj *log

func defaultPrintln(a ...interface{}) {
	fmt.Println(a...)
}

func defaultPrintf(fmtString string, a ...interface{}) {
	fmt.Printf(fmtString, a...)
}

func registerLogger() *zap.SugaredLogger {

	errorSyncer := getLogWriter("error")
	infoSyncer := getLogWriter("info")
	debugSyncer := getLogWriter("debug")
	ErrorLevel := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.ErrorLevel
	})
	DebugLevel := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	InfoLevel := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	encoder := getEncoder()
	error_core := zapcore.NewCore(encoder, errorSyncer, ErrorLevel)
	info_core := zapcore.NewCore(encoder, infoSyncer, InfoLevel)
	debug_core := zapcore.NewCore(encoder, debugSyncer, DebugLevel)

	logger := zap.New(zapcore.NewTee(error_core, info_core, debug_core), zap.AddCaller())
	return logger.Sugar()
}

// 初始化日志
func InitLogger() {
	logobj = new(log)

	if gin.Mode() != gin.DebugMode {
		logobj.logObj = registerLogger()
		{
			logobj.Info = logobj.logObj.Info
			logobj.Infof = logobj.logObj.Infof
			logobj.Errorf = logobj.logObj.Errorf
			logobj.Error = logobj.logObj.Error
			logobj.Debug = logobj.logObj.Debug
			logobj.Debugf = logobj.logObj.Debugf
		}
	} else {
		logobj.Info, logobj.Error, logobj.Debug = defaultPrintln, defaultPrintln, defaultPrintln
		logobj.Infof, logobj.Errorf, logobj.Debugf = defaultPrintf, defaultPrintf, defaultPrintf

	}

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(logType string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.LogConf.LogDir + logType + ".log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Info(args ...interface{}) {
	logobj.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logobj.Infof(format, args...)
}
func Debug(args ...interface{}) {
	logobj.Debug(args...)
}
func Debugf(format string, args ...interface{}) {
	logobj.Debugf(format, args...)
}

func Error(args ...interface{}) {
	logobj.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logobj.Errorf(format, args...)
}
