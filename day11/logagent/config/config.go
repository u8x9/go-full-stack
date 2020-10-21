package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	ip2 "github.com/u8x9/go-full-stack/day11/logagent/ip"
)

type kafkaConfig struct {
	Adderss     []string `toml:"address"`
	MaxChanSize int      `toml:"max_chan_size"`
}

type tailfConfig struct {
	Filename string `toml:"filename"`
}
type etcdConfig struct {
	Adderss       []string `toml:"address"`
	Timeout       int      `toml:"timeout"`
	CollectLogKey string   `toml:"collect_log_key"`
}

type LogAgentConfig struct {
	Kafka *kafkaConfig `toml:"kafka"`
	Etcd  *etcdConfig  `toml:"etcd"`
}

var Config LogAgentConfig

func Init() error {
	if _, err := toml.DecodeFile("config.toml", &Config); err != nil {
		return err
	}
	ip, err := ip2.GetOutboundIP()
	if err != nil {
		return err
	}
	Config.Etcd.CollectLogKey = fmt.Sprintf(Config.Etcd.CollectLogKey, ip)
	return nil
}
