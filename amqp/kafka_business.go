package amqp

import (
	"github.com/Shopify/sarama"
	"github.com/everywan/go-web-demo/config"
)

// 初始化配置(轨迹点)
func (k *kafkaHelper) init_tp() {
	// 从 .ini 文件读取配置
	*k = kafkaHelper{
		host:  config.ReadConfigByKey("./init.ini", "kafka", "host"),
		topic: config.ReadConfigByKey("./init.ini", "kafka", "topic_tp"),
	}
}

// 创建生产者(轨迹点)
func (k *kafkaHelper) createProducer_tp() {
	config := sarama.NewConfig()
	// 是否等待成功和失败后的响应, 只有RequireAcks设置不是NoReponse才生效.
	config.Producer.Return.Successes = true
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 设置使用的kafka版本
	// config.Version = sarama.V0_11_0_0
	k.createProducer(config)
}
