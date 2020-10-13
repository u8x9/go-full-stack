package mylogger

import (
	"fmt"
	"os"
	"time"
)

type FileLogger struct {
	Level       LogLevel
	fileName    string
    baseName string
	maxFileSize int64
	file        *os.File
	errFile     *os.File
}

func (f *FileLogger) mallocFile() {
	var filename string = f.baseName
	var errFilename = f.baseName + ".err"
	fileNeedSplit := false
	errFileNeedSplit := false
	if f.file != nil {
		fileNeedSplit, errFileNeedSplit, _ = f.needSplit()
		now := time.Now().Format("20160102150405")
		if fileNeedSplit {
			f.file.Close()
            newname := fmt.Sprintf("%s.%s", filename, now)
            os.Rename(filename, newname)
		}
		if errFileNeedSplit {
			f.errFile.Close()
            newname := fmt.Sprintf("%s.%s", errFilename, now)
            os.Rename(filename, newname)
		}
	}
	fm := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	fperm := os.FileMode(0644)
	if f.file == nil || fileNeedSplit {
		file, err := os.OpenFile(filename, fm, fperm)
		if err != nil {
			panic(err)
		}
		f.file = file
	}
	if f.errFile == nil || errFileNeedSplit {
		errFile, err := os.OpenFile(errFilename, fm, fperm)
		if err != nil {
			panic(err)
		}
		f.errFile = errFile
	}
}
func NewFileLogger(level LogLevel, fileName string, maxSize int64) *FileLogger {
	f := &FileLogger{
		Level:       level,
		fileName:    fileName,
		maxFileSize: maxSize,
		file:        nil,
		errFile:     nil,
        baseName: fileName,
	}
	f.mallocFile()
	return f
}
func (f *FileLogger) Close() {
	if f.errFile != nil {
		f.errFile.Close()
	}
	if f.file != nil {
		f.file.Close()
	}
}
func (f *FileLogger) enable(levevl LogLevel) bool {
	return f.Level <= levevl
}
func (f *FileLogger) log(level LogLevel, format string, a ...interface{}) {
	if !f.enable(level) {
		return
	}
	f.mallocFile()
	msg := fmt.Sprintf(format, a...)
	fnName, filename, line := getCallerInfo(3)
	data := newLogData(level, msg, fnName, filename, line).toString()
	fmt.Fprint(f.file, data)
	if f.errFile != nil && level >= ERROR {
		fmt.Fprint(f.errFile, data)
	}
}
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}
func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, format, a...)
}
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
func (f *FileLogger) needSplit() (bool, bool, error) {
	fileInfo, err := f.file.Stat()
	if err != nil {
		return false, false, err
	}
	errFileInfo, err := f.errFile.Stat()
	if err != nil {
		return false, false, err
	}
	return fileInfo.Size() > f.maxFileSize, errFileInfo.Size() > f.maxFileSize, nil
}
