package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// kafka 消费者

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Println("create consumer failed:", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Println("get Partitions failed:", err)
		return
	}
	fmt.Println("partitionList:", partitionList)

	for partition := range partitionList {
		// 针对每个分区创建一个对应的消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("start consumer failed on", partition, ":", err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition: %d, Offset: %d, Key: %v, Value: %v\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
}
