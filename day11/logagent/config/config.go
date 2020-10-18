package config

import "github.com/BurntSushi/toml"

type kafkaConfig struct {
    Adderss []string `toml:"address"`
    Topic   string `toml:"topic"`
}

type tailfConfig struct {
    Filename string `toml:"filename"`
}

type LogAgentConfig struct {
    Kafka *kafkaConfig `toml:"kafka"`
    Tailf *tailfConfig `toml:"tailf"`
}

var Config LogAgentConfig
func Init() error {
    if _,err:=toml.DecodeFile("config.toml", &Config);err!=nil {
        return err
    }
    return nil
}
