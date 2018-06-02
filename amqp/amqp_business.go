package amqp

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
)

// kafka测试, 以轨迹点为例
var gloableKafka_tp *kafkaHelper

// 初始化所需要的kafka消费者
func init() {
	gloableKafka_tp.init_tp()
	// 创建生产者
	gloableKafka_tp.createProducer_tp()
}

// 发送轨迹点
func Send_TrackPoint(bikeId string, curTime string, lon string, lat string) error {
	_map := map[string]string{
		"bikeId":  bikeId,
		"curTime": curTime,
		"lon":     lon,
		"lat":     lat,
	}
	value, err := json.Marshal(_map)
	if err == nil {
		return gloableKafka_tp.sendMsg(value)
	}
	return err
}

// 接收轨迹点
func Rece_TrackPoint() {
	gloableKafka_tp.receiveMsg(func(msg *sarama.ConsumerMessage) {
		fmt.Printf("收到消息: Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	})
}
