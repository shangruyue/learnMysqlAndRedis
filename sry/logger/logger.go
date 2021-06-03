package logger

import (
	//"log"
	"net/http"
	//"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

func InitSugarLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

//使用第三方包对日志进行分割归档
func getLogWriter() zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "./zaplog.log", //指定写入文件的位置
		MaxSize:    1,              //切割前日志文件的最大大小
		MaxBackups: 5,              // 保留旧文件的最大个数
		MaxAge:     30,             // 保留旧文件的最大天数
		Compress:   false,          //是否压缩和归档旧文件
	}
	return zapcore.AddSync(lumberjackLogger)
}

func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()) //以json格式打印日志
	//return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()) //不以json格式打印日志
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   //将时间改写可读形式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder //使用大写字母记录日志级别 还有颜色哎！
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func SetupSugarLogger() {
	InitSugarLogger()
	simpleHttpWithSugarZap("www.google.com")
	simpleHttpWithSugarZap("http://www.google.com")
}

func simpleHttpWithSugarZap(url string) {
	sugarLogger.Debugf("Trying to hit Get request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s", url)
	} else {
		sugarLogger.Info("Success! statusCode = ", resp.StatusCode)
	}
}

//使用zap做日志管理
/*var logger *zap.Logger

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func simpleHttpGetWithZap(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url",
			zap.String("url", url),
			zap.Error(err),
		)
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url),
		)
	}
}

func SetupLoggerWithZap() {
	InitLogger()
	defer logger.Sync()
	simpleHttpGetWithZap("www.google.com")
	simpleHttpGetWithZap("http://www.google.com")
}
*/

//使用内置的log包来实现log的管理
/*
func SetupLogger() {
	logFileLocation, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744) //将log内容输出到指定的文件中
	log.SetOutput(logFileLocation)

	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("status code for %s : %s", url, resp.Status)
	}
}
*/
