package main

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	Info
	WARNING
	ERROR
	FATAL
)

func parseloglevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return Info, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err

	}
}

func getlogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case Info:
		return "Info"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}

func getInfo(skip int) (funcName, fileName string, lineNO int) {
	pc, file, lineNO, ok := runtime.Caller(skip)
	if ok == false {
		fmt.Printf("runetime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	fileName = strings.Split(file, ".")[1]
	return
}

// logger 日志结构体
type Logger struct {
	Level LogLevel
}

// Newlog 构造函数
func Newlog(levelStr string) Logger {
	level, err := parseloglevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}

}
func (l Logger) enable(levellog LogLevel) bool {
	return levellog >= l.Level
}

func log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s: %s: %d] %s\n", now.Format("2006-01-02 15:04:05"), getlogString(lv), fileName, funcName, lineNo, msg)
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		log(DEBUG, msg)
	}
}

func (l Logger) Info(format string, a ...interface{}) {
	if l.enable(Info) {
		log(Info, format, a...)

	}
}
func (l Logger) Warning(format string, a ...interface{}) {
	if l.enable(WARNING) {
		log(WARNING, format, a...)
	}
}

func (l Logger) Error(format string, a ...interface{}) {
	if l.enable(ERROR) {
		log(ERROR, format, a...)

	}
}
func (l Logger) Fatal(format string, a ...interface{}) {
	if l.enable(FATAL) {
		log(FATAL, format, a...)

	}
}

func main() {
	log := Newlog("error")
	for {
		log.Debug("这是Debug日志")
		log.Info("这是Info日志")
		log.Warning("这是Warning日志")
		id := 1000
		name := "嘿嘿"
		log.Error("这是Error日志,%d,%s", id, name)
		log.Fatal("这是Fatal日志")
		time.Sleep(time.Second * 3)

	}
}
