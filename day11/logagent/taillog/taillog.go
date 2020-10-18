package taillog

import "github.com/hpcloud/tail"

// 专门从日志文件收集日志的模块

var (
	tails *tail.Tail
)

func Init(filename string) (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	tails, err = tail.TailFile(filename, config)
	return err
}
func Read() (ch <-chan *tail.Line) {
	return tails.Lines
}
