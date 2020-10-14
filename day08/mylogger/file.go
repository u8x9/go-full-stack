package mylogger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type FileLogger struct {
	level    LogLevel
	fileName string
	file     *os.File
	errFile  *os.File
	maxSize  int64
	ch       chan *logMessage
}

const extName = ".log"

func NewFileLogger(level LogLevel, fileName string, maxSize int64) *FileLogger {
	f := &FileLogger{
		level:    level,
		fileName: fileName,
		file:     nil,
		errFile:  nil,
		maxSize:  maxSize,
		ch:       make(chan *logMessage, 1024),
	}
	fn, efn := f.getFileName(false, false)
	f.mallocFile(fn, efn)
	for i := 0; i < 1; i++ {
		go f.background()
	}
	return f
}

func (f *FileLogger) getFileName(isSplitFile, isSplitErrFile bool) (filename, errFilename string) {
	filename = f.fileName
	if strings.HasSuffix(strings.ToLower(filename), extName) {
		filename = strings.TrimSuffix(filename, extName)
	}
	errFilename = filename + ".err"
	ts := time.Now().Format("20060102150405")
	if isSplitFile {
		filename = fmt.Sprintf("%s.%s", filename, ts)
	}
	if isSplitErrFile {
		errFilename = fmt.Sprintf("%s.%s", errFilename, ts)
	}

	filename += extName
	errFilename += extName
	return filename, errFilename
}

func (f *FileLogger) mallocFile(fileName, errFileName string) error {
	flag := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	mode := os.FileMode(0644)
	if len(fileName) > 0 {
		f.closeFile()
		file, err := os.OpenFile(fileName, flag, mode)
		if err != nil {
			return err
		}
		f.file = file
	}
	if len(errFileName) > 0 {
		f.closeErrFile()
		errFile, err := os.OpenFile(errFileName, flag, mode)
		if err != nil {
			return err
		}
		f.errFile = errFile
	}
	return nil
}

func (f *FileLogger) closeFile() {
	if f.file != nil {
		f.file.Close()
	}
}
func (f *FileLogger) closeErrFile() {
	if f.errFile != nil {
		f.errFile.Close()
	}
}
func (f *FileLogger) Close() {
	f.closeFile()
	f.closeErrFile()
}
func (f *FileLogger) needSplit() (bool, bool, error) {
	info, err := f.file.Stat()
	if err != nil {
		return false, false, err
	}
	errInfo, err := f.errFile.Stat()
	if err != nil {
		return false, false, err
	}
	return info.Size() >= f.maxSize, errInfo.Size() >= f.maxSize, nil
}

func (f *FileLogger) getLevel() LogLevel {
	return f.level
}
func (f *FileLogger) enable(level LogLevel) bool {
	return enable(f, level)
}
func (f *FileLogger) background() {
	for {
		select {
		case logMsg := <-f.ch:
			isNeedSplitFile, isNeedSplitErrFile, _ := f.needSplit()
			fn, efn := f.getFileName(isNeedSplitFile, isNeedSplitErrFile)
			if !isNeedSplitFile {
				fn = ""
			}
			if !isNeedSplitErrFile {
				efn = ""
			}
			f.mallocFile(fn, efn)
			if _, err := fmt.Fprint(f.file, logMsg.toString()); err != nil {
				fmt.Println(err)
			}
			if logMsg.level >= ERROR {
				fmt.Fprint(f.errFile, logMsg.toString())
			}

		default:
			time.Sleep(time.Millisecond * 300)
		}
	}

}
func (f *FileLogger) log(level LogLevel, format string, a ...interface{}) {
	if !f.enable(level) {
		return
	}

	msg := fmt.Sprintf(format, a...)
	fnName, fileName, line := getCallInfo(3)
	logMsg := newLogMessage(msg, fnName, fileName, level, line)

	select {
	case f.ch <- logMsg:
		{
		}
	default:
		{
		}
	}
}
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}
