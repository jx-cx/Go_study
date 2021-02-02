package main

import (
	"errors"
	"fmt"
	"os"
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
type ConsoleLogger struct {
	Level LogLevel
}

func (f *Fileogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...) // 格式化赋值给msg
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Fprint(f.fileObj, "[%s] [%s] [%s: %s: %d] %s\n", now.Format("2006-01-02 15:04:05"), getlogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			fmt.Fprint(f.errfileObj, "[%s] [%s] [%s: %s: %d] %s\n", now.Format("2006-01-02 15:04:05"), getlogString(lv), fileName, funcName, lineNo, msg)
		}

	}
}
func (f *Fileogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

func (f *Fileogger) Info(format string, a ...interface{}) {

	f.log(Info, format, a...)

}
func (f *Fileogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)

}

func (f *Fileogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)

}
func (f *Fileogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)

}

// Newlog 构造函数
func Newlog(levelStr string) ConsoleLogger {
	level, err := parseloglevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}

}

// 往文件里面写日志
type Fileogger struct {
	level       LogLevel
	filepath    string // 日志保存的路径
	filename    string //日志保存的名称
	fileObj     *os.File
	errfileObj  *os.File
	maxFileSize int64
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *Fileogger {
	LogLevel, err := parseloglevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &Fileogger{
		level:       LogLevel,
		filepath:    fp,
		filename:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initfile()
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *Fileogger) enable(levellog LogLevel) bool {
	return levellog >= f.level
}

func (f *Fileogger) initfile() error {
	fullFileName := path.Join(f.filepath, f.filename)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed ,err:%v", err)
		return err
	}
	errfileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed ,err:%v", err)
		return err
	}
	f.fileObj = fileObj
	f.fileObj = errfileObj
	return nil
}

//关闭文件
func (f *Fileogger) Close() {
	f.fileObj.Close()
	f.errfileObj.Close()
}

func main() {
	log := NewFileLogger("info", "/", "哈哈", 10*1024*1024)
	for {
		log.Debug("这是Debug日志")
		log.Info("这是Info日志")
		log.Warning("这是Warning日志")
		id := 1000
		name := "嘿 嘿"
		log.Error("这是Error日志,%d,%s", id, name)
		log.Fatal("这是Fatal日志")
		time.Sleep(time.Second * 3)

	}
}
