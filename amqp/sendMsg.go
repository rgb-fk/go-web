package amqp

import (
	"encoding/json"
)

var gloableKafka_tp kafkaHelper

func init() {
	// 轨迹点 初始化
	gloableKafka_tp.init_tp()
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

// 接收轨迹点_测试
func Rece_TrackPoint() {
	gloableKafka_tp.receiveMsg_tp()
}
