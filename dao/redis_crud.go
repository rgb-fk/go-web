package dao

import (
	"fmt"

	"github.com/everywan/go-web/config"
	"github.com/garyburd/redigo/redis"
)

var gloableRedisHelper redisHelper

type redisHelper struct {
	network      string
	host         string
	password     string
	database     int
	connTimeout  int
	readTimeout  int
	writeTimeout int
}

func (r redisHelper) init() {
	// 读取配置文件, 初始化 gloableRedisHelper
	r = redisHelper{
		network:      config.ReadConfigByKey("./init.ini", "redis", "network"),
		host:         config.ReadConfigByKey("./init.ini", "redis", "host"),
		password:     config.ReadConfigByKey("./init.ini", "redis", "password"),
		database:     int(config.ReadConfigByKeyInt("./init.ini", "redis", "database")),
		connTimeout:  int(config.ReadConfigByKeyInt("./init.ini", "redis", "connTimeout")),
		readTimeout:  int(config.ReadConfigByKeyInt("./init.ini", "redis", "readTimeout")),
		writeTimeout: int(config.ReadConfigByKeyInt("./init.ini", "redis", "writeTimeout")),
	}
	gloableRedisHelper = r
}

// --------------------------------------------元数据------------------------------------------------------------------
func (r *redisHelper) getConn() (redis.Conn, error) {
	conn, err := redis.Dial(r.network, r.host,
		redis.DialPassword(r.password), redis.DialDatabase(r.database))
	return conn, err
}
func (r *redisHelper) closeConn(conn redis.Conn) error {
	err := conn.Close()
	return err
}

// ----------------------------------------------业务实现--------------------------------------------------------

// 获取TODO
func (r *redisHelper) getToDo_EI() (string, error) {
	conn, err := r.getConn()
	if err == nil {
		defer r.closeConn(conn)
		todos, err := redis.String(conn.Do("HGET", TODO, "EI"))
		return fmt.Sprint(todos), err
	}
	return "", err
}
func (r *redisHelper) getToDo_E() (string, error) {
	conn, err := r.getConn()
	if err == nil {
		defer r.closeConn(conn)
		todos, err := redis.String(conn.Do("HGET", TODO, "E"))
		return fmt.Sprint(todos), err
	}
	return "", err
}
func (r *redisHelper) getToDo_I() (string, error) {
	conn, err := r.getConn()
	if err == nil {
		defer r.closeConn(conn)
		todos, err := redis.String(conn.Do("HGET", TODO, "I"))
		return fmt.Sprint(todos), err
	}
	return "", err
}
func (r *redisHelper) getToDo_Other() (string, error) {
	conn, err := r.getConn()
	if err == nil {
		defer r.closeConn(conn)
		todos, err := redis.String(conn.Do("HGET", TODO, "Other"))
		return fmt.Sprint(todos), err
	}
	return "", err
}

// 添加一个TODO项
func (r *redisHelper) SetToDo(msg string, key string) error {
	conn, err := r.getConn()
	if err == nil {
		defer r.closeConn(conn)
		_, err = conn.Do("hset", TODO, key, msg)
	}
	return err
}
