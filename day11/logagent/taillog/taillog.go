package taillog

import (
	"context"
	"fmt"

	"github.com/hpcloud/tail"
	"github.com/u8x9/go-full-stack/day11/logagent/kafka"
)

// 专门从日志文件收集日志的模块

// TailTask 一个日志收集的任务
type TailTask struct {
	Filename string
	Topic    string
	Instance *tail.Tail
	ctx      context.Context
	cancel   context.CancelFunc
}

func NewTailTask(filename, topic string) *TailTask {
	ctx, cancel := context.WithCancel(context.Background())
	tt := &TailTask{
		Filename: filename,
		Topic:    topic,
		Instance: nil,
		ctx:      ctx,
		cancel:   cancel,
	}
	tt.init()
	return tt
}

func (tt *TailTask) init() {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	var err error
	tt.Instance, err = tail.TailFile(tt.Filename, config)
	if err != nil {
		fmt.Println("TailFile failed:", err)
		return
	}
	go tt.Run()
}

func (tt *TailTask) Run() {
	for {
		select {
		case <-tt.ctx.Done():
			fmt.Println("我的工作完成了，撒油那拉～", genMapKey(tt.Filename, tt.Topic))
			return
		case line := <-tt.Instance.Lines:
			//kafka.SendMessageToTopic(tt.Topic, line.Text)
			// 先把日志数据发到一个channel中
			// kafka 那个包中有单独的goroutine去取日志数据发到kafka服务器
			kafka.SendToChan(tt.Topic, line.Text)
		}
	}
}
