package mylogger

import (
	"fmt"
)

type ConsoleLogger struct {
	level LogLevel
}

func NewConsoleLogger(level LogLevel) *ConsoleLogger {
	return &ConsoleLogger{
		level: level,
	}
}

func (c *ConsoleLogger) getLevel() LogLevel {
	return c.level
}
func (c *ConsoleLogger) enable(level LogLevel) bool {
	return enable(c, level)
}
func (c *ConsoleLogger) log(level LogLevel, format string, a ...interface{}) {
	if !c.enable(level) {
		return
	}
	msg := fmt.Sprintf(format, a...)
	fnName, fileName, line := getCallInfo(3)
	logMsg := newLogMessage(msg, fnName, fileName, level, line)
	fmt.Print(logMsg.toString())
}
func (c *ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}
func (c *ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}
func (c *ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}
