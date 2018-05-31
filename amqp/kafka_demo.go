package amqp

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Shopify/sarama"
)

func ProducerTest() error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewSyncProducer([]string{"192.168.1.174:9092"}, config)
	if err == nil {
		defer producer.Close()

		msg := &sarama.ProducerMessage{
			Topic:     "trackPoint",
			Partition: int32(-1),
			Key:       sarama.StringEncoder("key"),
		}

		var value string
		for {
			// 生产消息
			inputReader := bufio.NewReader(os.Stdin)
			value, err = inputReader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			value = strings.Replace(value, "\n", "", -1)
			msg.Value = sarama.ByteEncoder(value)
			paritition, offset, err := producer.SendMessage(msg)

			if err != nil {
				fmt.Println("Send Message Fail")
			}

			fmt.Printf("Partion = %d, offset = %d\n", paritition, offset)
		}
	} else {
		return err
	}
}

func ConsumerTest() {
	consumer, err := sarama.NewConsumer([]string{"192.168.1.174:9092"}, nil)

	if err != nil {
		panic(err)
	}

	partitionList, err := consumer.Partitions("trackPoint")

	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("trackPoint", int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}

		defer pc.AsyncClose()

		wg.Add(1)

		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("收到消息Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
		wg.Wait()
		consumer.Close()
	}
}
