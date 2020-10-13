package mylogger

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

type LogLevel uint8

const (
	DEBUG LogLevel = iota
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

type logData struct {
	date     string
	level    string
	msg      string
	fnName   string
	fileName string
	line     int
}

func newLogData(level LogLevel, msg, fnName, fileName string, line int) *logData {
	return &logData{
        date:     time.Now().Format("2006/01/02 15:04:05"),
		level:    logLevel2str(level),
		msg:      msg,
		fnName:   fnName,
		fileName: fileName,
		line:     line,
	}
}
func (ld *logData) toString() string {
	return fmt.Sprintf("[%s] [%s] %s -%s:%s@%d\n", ld.date, ld.level, ld.msg, ld.fnName, ld.fileName, ld.line)
}

func logLevel2str(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "LOG"
}

type Logger interface {
	enable(level LogLevel) bool
	log(level LogLevel, format string, a ...interface{})
	Debug(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

func getCallerInfo(skip int) (funcName, fileName string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}
