package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type LogConfigEntry struct {
	Filename string `json:"filename"`
	Topic    string `json:"topic"`
}

var cli *clientv3.Client

// Init 初始化etcd
func Init(address []string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   address,
		DialTimeout: timeout,
	})
	if err != nil {
		return err
	}
	return nil
}

// Close 关闭
func Close() error {
	return cli.Close()
}

// Watch 监控
func Watch(key string, newCfgChan chan<- []*LogConfigEntry) {
	ch := cli.Watch(context.Background(), key)
	for wresp := range ch {
		for _, evt := range wresp.Events {
			// 通知 taillog.tskMgr
			var newCfg []*LogConfigEntry
			if evt.Type !=clientv3.EventTypeDelete {
				if err := json.Unmarshal(evt.Kv.Value, &newCfg); err != nil {
					fmt.Println("unmarshal json failed:", err)
					continue
				}
			}
			newCfgChan <- newCfg
		}
	}
}

func Get(key string) (logConfigEnties []*LogConfigEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		return nil, err
	}

	for _, ev := range resp.Kvs {
		if err = json.Unmarshal(ev.Value, &logConfigEnties); err != nil {
			return nil, err
		}
	}

	return logConfigEnties, nil
}
func Put(key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := cli.Put(ctx, key, value)
	cancel()
	return err
}
