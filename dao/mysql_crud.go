package dao

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/everywan/go-web/config"
	_ "github.com/go-sql-driver/mysql"
)

var gloableMysqlHelper mysqlHelper

type mysqlConfig struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type mysqlHelper struct {
	network  string
	host     string
	userName string
	password string
	database string
	db       *sql.DB
}

func (m mysqlHelper) init() {
	m = mysqlHelper{
		network:  config.ReadConfigByKey("./init.ini", "mysql", "network"),
		host:     config.ReadConfigByKey("./init.ini", "mysql", "host"),
		userName: config.ReadConfigByKey("./init.ini", "mysql", "userName"),
		password: config.ReadConfigByKey("./init.ini", "mysql", "password"),
		database: config.ReadConfigByKey("./init.ini", "mysql", "database"),
	}
	url := fmt.Sprintf("%s:%s@%s(%s)/%s", m.userName, m.password, m.network, m.host, m.database)
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic("无法链接数据库, dburl: " + url)
	}
	mc := mysqlConfig{
		MaxOpenConns:    1000,
		MaxIdleConns:    10,
		ConnMaxLifetime: time.Second * 10,
	}
	db.SetMaxOpenConns(mc.MaxOpenConns)
	db.SetMaxIdleConns(mc.MaxIdleConns)
	db.SetConnMaxLifetime(mc.ConnMaxLifetime)

	m.db = db
	gloableMysqlHelper = m
}

func (m *mysqlHelper) Demo() {
	rows, err := m.db.Query("select id,orderNum from kh_orders limit 2")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var orderNum string
		if err := rows.Scan(&id, &orderNum); err != nil {
			fmt.Println("error in rows")
			fmt.Println(err)
		}
		fmt.Println(id, orderNum)
	}
}
