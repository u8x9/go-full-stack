package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/u8x9/go-full-stack/day11/logagent/config"
	"github.com/u8x9/go-full-stack/day11/logagent/etcd"
	"github.com/u8x9/go-full-stack/day11/logagent/kafka"
	"github.com/u8x9/go-full-stack/day11/logagent/taillog"
)

func main() {
	// 0. 读取配置
	if err := config.Init(); err != nil {
		panic(err)
	}
	// 1. 初始化kafka连接
	if err := kafka.Init(config.Config.Kafka.Adderss, config.Config.Kafka.MaxChanSize); err != nil {
		fmt.Println("Init kafka failed:", err)
		return
	}
	fmt.Println("Kafka init success.")
	defer kafka.Close()
	// 2. 初始化etcd
	etcdTimeout := time.Duration(config.Config.Etcd.Timeout) * time.Second
	if err := etcd.Init(config.Config.Etcd.Adderss, etcdTimeout); err != nil {
		fmt.Println("Init etcd failed:", err)
		return
	}
	fmt.Println("Etcd init success.")
	defer etcd.Close()
	// 2.1 从etcd中获取日志收集项的配置信息
	fmt.Println(config.Config.Etcd.CollectLogKey)
	logCfgEnties, err := etcd.Get(config.Config.Etcd.CollectLogKey)
	if err != nil {
		fmt.Println("Get config from etcd failed:", err)
		return
	}
	// 2.2 派一个哨兵去监视配置信息的变动，及时通知LogAgent实现热加载
	taillog.Init(logCfgEnties)
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.Watch(config.Config.Etcd.CollectLogKey, taillog.NewConfChan())
	wg.Wait()
	// 3. 打开日志文件准备收集日志
	//if err := taillog.Init(config.Config.Tailf.Filename); err != nil {
	//panic(err)
	//}
	// 4. 运行
	//run(config.Config.Kafka.Topic)
}
