package mylogger

import (
	"fmt"
)

type ConsoleLogger struct {
    Level LogLevel
}

func NewConsoleLogger(level LogLevel) *ConsoleLogger {
    return &ConsoleLogger {
        Level: level,
    }
}

func (c *ConsoleLogger) enable(level LogLevel) bool {
    return c.Level <= level
}
func (c *ConsoleLogger) log(level LogLevel, format string, a ...interface{}) {
    if !c.enable(level) {
        return
    }
    msg:=fmt.Sprintf(format, a...)
    fnName, filename, line :=getCallerInfo(3)
	data := newLogData(level, msg, fnName, filename, line).toString()
    fmt.Printf(data)
}

func (c *ConsoleLogger) Debug(format string, a ...interface{}) {
    c.log(DEBUG, format, a...)
}
func (c *ConsoleLogger) Trace(format string, a ...interface{}) {
    c.log(TRACE, format, a...)
}
func (c *ConsoleLogger) Info(format string, a ...interface{}) {
    c.log(INFO, format, a...)
}
func (c *ConsoleLogger) Warning(format string, a ...interface{}) {
    c.log(WARNING, format, a...)
}
func (c *ConsoleLogger) Error(format string, a ...interface{}) {
    c.log(ERROR, format, a...)
}
func (c *ConsoleLogger) Fatal(format string, a ...interface{}) {
    c.log(FATAL, format, a...)
}
