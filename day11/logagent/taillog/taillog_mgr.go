package taillog

import (
	"fmt"
	"github.com/u8x9/go-full-stack/day11/logagent/etcd"
	"time"
)

type tailLogMgr struct {
	logEntry    []*etcd.LogConfigEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogConfigEntry
}

var tskMgr *tailLogMgr

func Init(logEntry []*etcd.LogConfigEntry) {
	tskMgr = &tailLogMgr{
		logEntry:    logEntry,
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogConfigEntry), // 无缓冲区的channel
	}
	for _, le := range logEntry {
		tailTask := NewTailTask(le.Filename, le.Topic)
		key := genMapKey(tailTask.Filename, tailTask.Topic)
		tskMgr.tskMap[key] = tailTask
	}
	go tskMgr.run()
}

func genMapKey(filename, topic string) string {
	return fmt.Sprintf("%s%s", filename, topic)
}

// run 监听自己的 newConfChan，有了新的配置就进行对应的处理
func (tlm *tailLogMgr) run() {
	for {
		select {
		case newCfg := <-tlm.newConfChan:
			for _, cfg := range newCfg {
				key := genMapKey(cfg.Filename, cfg.Topic)
				if _, ok := tlm.tskMap[key]; ok {
					continue
				} else {
					tlm.tskMap[key] = NewTailTask(cfg.Filename, cfg.Topic)
				}
			}

			for _, c1 := range tlm.logEntry {
				isDelete := true
				for _, c2 := range newCfg {
					if c1.Topic == c2.Topic && c1.Filename == c2.Filename {
						isDelete = false
					}
				}
				if isDelete {
					key := genMapKey(c1.Filename, c1.Topic)
					tlm.tskMap[key].cancel()
				}
			}
			fmt.Println("新的配置来了：", newCfg)
		default:
			time.Sleep(time.Second)
		}
	}
}

func NewConfChan() chan<- []*etcd.LogConfigEntry {
	return tskMgr.newConfChan
}
