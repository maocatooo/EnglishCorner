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
	logObj *zap.SugaredLogger
	Funcf  func(string, ...interface{})
	Func   func(...interface{})
}

var log_info *log
var log_error *log
var log_debug *log

func main() {
	InitLogger()
	defer log_info.logObj.Sync()
	defer log_error.logObj.Sync()
	defer log_debug.logObj.Sync()
}
func defaultPrintln(a ...interface{}) {
	fmt.Println(a...)
}
func defaultPrintf(fmtString string, a ...interface{}) {
	fmt.Printf(fmtString, a...)
}
func registerLogger(logType string) *zap.SugaredLogger {

	writeSyncer := getLogWriter(logType)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}
func InitLogger() {
	log_info = new(log)
	log_debug = new(log)
	log_error = new(log)
	if gin.Mode() != gin.DebugMode {
		log_info.logObj = registerLogger("info")
		{
			log_info.Func = log_info.logObj.Info
			log_info.Funcf = log_info.logObj.Infof
		}
		log_debug.logObj = registerLogger("debug")
		{
			log_debug.Func = log_debug.logObj.Debug
			log_debug.Funcf = log_debug.logObj.Debugf
		}
		log_error.logObj = registerLogger("error")
		{
			log_error.Func = log_error.logObj.Error
			log_error.Funcf = log_error.logObj.Errorf
		}
	} else {
		{
			log_info.Func = defaultPrintln
			log_info.Funcf = defaultPrintf
			log_error.Func = defaultPrintln
			log_error.Funcf = defaultPrintf
			log_debug.Func = defaultPrintln
			log_debug.Funcf = defaultPrintf
		}
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
	log_info.Func(args...)
}

func Infof(format string, args ...interface{}) {
	log_info.Funcf(format, args...)
}
func Debug(args ...interface{}) {
	log_debug.Func(args...)
}
func Debugf(format string, args ...interface{}) {
	log_debug.Funcf(format, args...)
}

func Error(args ...interface{}) {
	log_error.Func(args...)
}

func Errorf(format string, args ...interface{}) {
	log_error.Funcf(format, args...)
}
