package amqp

import (
	"sync"

	"github.com/Shopify/sarama"
)

var (
	msg           *sarama.ProducerMessage
	producer      sarama.SyncProducer
	wg            sync.WaitGroup
	consumer      sarama.Consumer
	partitionList []int32
)

type kafkaHelper struct {
	host  string
	topic string
}

// 新建一个异步生产者
func (k *kafkaHelper) createProducer(config *sarama.Config) {
	var err error
	producer, err = sarama.NewSyncProducer([]string{k.host}, config)
	if err != nil {
		panic(err)
	}
	// 发送的消息, 设置topic, 分区, key
	msg = &sarama.ProducerMessage{
		Topic:     k.topic,
		Partition: int32(-1),
		Key:       sarama.StringEncoder("key"),
	}
}

// 发送消息, []byte 可以隐式转换为 sarama.ByteEncoder
// map等复杂结构可以使用 `json.Marshal(map)`` 序列化为[]byte
func (k *kafkaHelper) sendMsg(value sarama.ByteEncoder) error {
	msg.Value = value
	_, _, err := producer.SendMessage(msg)
	return err
}

// 接收消息, 参数为对于每条msg的处理方法
func (k *kafkaHelper) receiveMsg(handle func(*sarama.ConsumerMessage)) {
	var err error

	// 创建消费者
	if consumer, err = sarama.NewConsumer([]string{k.host}, nil); err != nil {
		panic(err)
	}
	// 关闭消费者
	defer consumer.Close()

	// 获取分区信息
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
			// 从每个分区拉取消息
			for msg := range pc.Messages() {
				handle(msg)
			}
		}(pc)
		wg.Wait()
	}
}
