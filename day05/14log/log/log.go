package log

import (
	"fmt"
	"io"
)

const (
	DEBUG = iota
	INFO
	WARNNING
	ERROR
)

type Logger struct {
	w io.Writer
}

func NewLogger(w io.Writer) *Logger {
	return &Logger{
		w: w,
	}
}

func (l *Logger) log(level int, msg string) (int, error) {
	m := make(map[int]string, 4)
	m[DEBUG] = "DEBUG"
	m[INFO] = "INFO"
	m[WARNNING] = "WARN"
	m[ERROR] = "ERROR"
	levelStr := "LOG"
	if v, ok := m[level]; ok {
		levelStr = v
	}
	buf := []byte(fmt.Sprintf("[%s] %s\n", levelStr, msg))
	return l.w.Write([]byte(buf))
}

func (l *Logger) Debug(msg string) (int, error) {
	return l.log(DEBUG, msg)
}

func (l *Logger) Info(msg string) (int, error) {
	return l.log(INFO, msg)
}

func (l *Logger) Warn(msg string) (int, error) {
	return l.log(WARNNING, msg)
}

func (l *Logger) Error(msg string) (int, error) {
	return l.log(ERROR, msg)
}

