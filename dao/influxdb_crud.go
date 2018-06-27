package dao

import (
	"fmt"

	"github.com/everywan/go-web/config"
)

var gloableInfluxdbHelper influxdbHelper

type influxdbHelper struct {
	host     string
	database string
	user     string
	password string
	url      string
}

func (influx influxdbHelper) init() {
	// 读取配置文件
	influx = influxdbHelper{
		host:     config.ReadConfigByKey("./init.ini", "influxdb", "host"),
		database: config.ReadConfigByKey("./init.ini", "influxdb", "database"),
		user:     config.ReadConfigByKey("./init.ini", "influxdb", "user"),
		password: config.ReadConfigByKey("./init.ini", "influxdb", "password"),
	}
	influx.url = fmt.Sprintf("http://%s/write?db=%s&u=%s&p=%s", influx.host, influx.database, influx.user, influx.password)
	gloableInfluxdbHelper = influxdbHelper{}
	gloableInfluxdbHelper = influx
}
