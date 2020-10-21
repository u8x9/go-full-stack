package kafka

import (
	"time"

	"github.com/Shopify/sarama"
)

// 专门将日志写入kafka的模块

type logData struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer
	logDataChan chan *logData
)

// Init 初始化 kafka 连接
func Init(address []string, chanMaxSize int) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	var err error
	if client, err = sarama.NewSyncProducer(address, config); err != nil {
		return err
	}
	logDataChan = make(chan *logData, chanMaxSize)
	go sendMessage()
	return nil
}

// SendToChan 给外部暴露一个函数，该函数只把日志数据发送到一个内部的channel 中
func SendToChan(topic, data string) {
	logDataChan <- &logData{
		topic: topic,
		data:  data,
	}
}

// 真正发送消息到指定kafka
func sendMessage() {
	for {
		select {
		case ld := <-logDataChan:
			msg := new(sarama.ProducerMessage)
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			client.SendMessage(msg)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}

// Close 关闭连接
func Close() error {
	return client.Close()
}
