package amqp

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/everywan/go-web-demo/config"
)

type kafkaHelper struct {
	host  string
	topic string
}

var (
	producer      sarama.SyncProducer
	msg           *sarama.ProducerMessage
	wg            sync.WaitGroup
	consumer      sarama.Consumer
	partitionList []int32
)

// 初始化配置(轨迹点)
func (k *kafkaHelper) init_tp() {
	*k = kafkaHelper{
		host:  config.ReadConfigByKey("./init.ini", "kafka", "host"),
		topic: config.ReadConfigByKey("./init.ini", "kafka", "topic_tp"),
	}
}

// 创建生产者(轨迹点)
func (k *kafkaHelper) createProducer_tp() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	var err error
	producer, err = sarama.NewSyncProducer([]string{k.host}, config)
	if err != nil {
		panic(err)
	}
	msg = &sarama.ProducerMessage{
		Topic:     k.topic,
		Partition: int32(-1),
		Key:       sarama.StringEncoder("key"),
	}
}

//发送消息
func (k *kafkaHelper) sendMsg(value sarama.ByteEncoder) error {
	// msg.Value = sarama.ByteEncoder(value)
	// paritition, offset, err := k.producer.SendMessage(k.msg)
	// fmt.Println(value)
	msg.Value = value
	_, _, err := producer.SendMessage(msg)
	return err
}

// 接收消息(轨迹点)
func (k *kafkaHelper) receiveMsg_tp() {
	var err error
	if consumer, err = sarama.NewConsumer([]string{k.host}, nil); err != nil {
		panic(err)
	}

	if partitionList, err = consumer.Partitions(k.topic); err != nil {
		panic(err)
	}
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(k.topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}

		defer pc.AsyncClose()
		wg.Add(1)

		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				// 收到消息后的处理
				fmt.Printf("收到消息: Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
		wg.Wait()
		consumer.Close()
	}
}
