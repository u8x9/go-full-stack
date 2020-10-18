package kafka

import (
	"github.com/Shopify/sarama"
)

// 专门将日志写入kafka的模块

var (
	client       sarama.SyncProducer
	locatAddress = []string{"127.0.0.1:9092"}
)

const DEFAULT_TOPIC = "web_log"

// Init 初始化 kafka 连接
func Init(address []string) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	var err error
	if client, err = sarama.NewSyncProducer(address, config); err != nil {
		return err
	}
	return nil
}

// InitDefaultLocalSingle 初始化本地单节点kafka
func InitDefaultLocalSingle() error {
	return Init(locatAddress)
}

// SendMessageToTopic 发送消息到指定topic
func SendMessageToTopic(topic, message string) (pid int32, offset int64, err error) {
	msg := new(sarama.ProducerMessage)
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(message)
	pid, offset, err = client.SendMessage(msg)
	return
}

// SendMessage 发送消息到默认topic
func SendMessage(message string) (pid int32, offset int64, err error) {
	return SendMessageToTopic(DEFAULT_TOPIC, message)
}

// Close 关闭连接
func Close() error {
	return client.Close()
}
