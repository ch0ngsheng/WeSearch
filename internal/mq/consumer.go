package mq

import (
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

type Consumer interface {
	Consume(group, topic string, fn func(msg []byte) error)
	io.Closer
}

type kafkaConsumer struct {
	client sarama.Client
}

func (k kafkaConsumer) Close() error {
	return k.client.Close()
}

func (k kafkaConsumer) Consume(group, topic string, fn func(msg []byte) error) {
	var partitions []int32
	var err error

	RetryWith("get-partitions", time.Second*2, 60, func() error {
		partitions, err = k.client.Partitions(topic)
		if err != nil {
			log.Printf("fail to list partitions, %v", err)
		}
		return err
	})

	var wg sync.WaitGroup
	wg.Add(len(partitions))
	for _, partitionId := range partitions {
		go k.consumeByPartition(&wg, group, topic, partitionId, fn)
	}

	wg.Wait()
}

func (k kafkaConsumer) consumeByPartition(wg *sync.WaitGroup, group, topic string, partitionId int32, fn func(msg []byte) error) {
	defer wg.Done()

	var err error
	var offsetManager sarama.OffsetManager

	RetryWith("offset-manager", time.Second*2, 60, func() error {
		offsetManager, err = sarama.NewOffsetManagerFromClient(group, k.client)
		if err != nil {
			log.Println("NewOffsetManagerFromClient err:", err)
		}
		return err
	})
	defer offsetManager.Close()

	// 分区的 offset 也是分别管理的
	var partitionOffsetManager sarama.PartitionOffsetManager
	RetryWith("partition-offset-manager", time.Second*2, 60, func() error {
		partitionOffsetManager, err = offsetManager.ManagePartition(topic, partitionId)
		if err != nil {
			log.Println("ManagePartition err:", err)
		}
		return err
	})
	defer partitionOffsetManager.Close()

	// 程序结束后 commit 一次，防止自动提交间隔之间的信息被丢掉
	defer offsetManager.Commit()

	var consumer sarama.Consumer
	var nextOffset int64
	var partitionConsumer sarama.PartitionConsumer

	RetryWith("get-consumer", time.Second*2, 60, func() error {
		consumer, err = sarama.NewConsumerFromClient(k.client)
		if err != nil {
			log.Println("NewConsumerFromClient err:", err)
			return err
		}

		// 根据 kafka 中记录的上次消费的 offset 开始+1的位置接着消费
		nextOffset, _ = partitionOffsetManager.NextOffset()

		partitionConsumer, err = consumer.ConsumePartition(topic, partitionId, nextOffset)
		if err != nil {
			log.Println("ConsumePartition err:", err)
			return err
		}

		return nil
	})
	defer partitionConsumer.Close()
	fmt.Println("nextOffset:", nextOffset)

	for message := range partitionConsumer.Messages() {
		value := string(message.Value)
		log.Printf("[Consumer] group %s, partitionId: %d, offset:%d, value: %s\n", group, message.Partition, message.Offset, value)

		// 业务回调
		if err = fn(message.Value); err != nil {
			log.Printf("fail to process message: %s, error: %v", value, err)
		}

		// 每次消费后都更新一次 offset。这里更新的只是内存中的值，需要 commit 之后才能提交到 kafka
		partitionOffsetManager.MarkOffset(message.Offset+1, "modified metadata")
	}
}
