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
	INFO
	ERROR
)

type Logger interface {
	getLevel() LogLevel
	enable(LogLevel) bool
	log(LogLevel, string, ...interface{})
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Error(string, ...interface{})
}

type logMessage struct {
	msg      string
	date     string
	level    LogLevel
	fnName   string
	fileName string
	line     int
}

func newLogMessage(msg, fnName, fileName string, level LogLevel, line int) *logMessage {
	return &logMessage{
		msg:      msg,
		date:     time.Now().Format("2006/01/02 15:04:05"),
		level:    level,
		fnName:   fnName,
		fileName: fileName,
		line:     line,
	}
}
func (lm *logMessage) toString() string {
	return fmt.Sprintf("[%s] [%s] %s -%s:%s@%d\n", lm.date, lm.level.toString(), lm.msg, lm.fnName, lm.fileName, lm.line)
}

func enable(logger Logger, level LogLevel) bool {
	return logger.getLevel() <= level
}

func (l LogLevel) toString() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case ERROR:
		return "ERROR"
	}
	return "LOG"
}

func getCallInfo(skip int) (fnName, fileName string, line int) {
	if pc, file, ln, ok := runtime.Caller(skip); ok {
		fnName = runtime.FuncForPC(pc).Name()
		fileName = path.Base(file)
		line = ln
	}
	return
}
