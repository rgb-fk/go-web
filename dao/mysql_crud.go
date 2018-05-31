package dao

import (
	"database/sql"
	"fmt"

	"github.com/everywan/go-web-demo/config"
	_ "github.com/go-sql-driver/mysql"
)

var gloableMysqlHelper mysqlHelper

type mysqlHelper struct {
	network  string
	host     string
	userName string
	password string
	database string
	// connTimeout  int
	// readTimeout  int
	// writeTimeout int
	url string
}

func (m mysqlHelper) init() {
	m = mysqlHelper{
		network:  config.ReadConfigByKey("./init.ini", "mysql", "network"),
		host:     config.ReadConfigByKey("./init.ini", "mysql", "host"),
		userName: config.ReadConfigByKey("./init.ini", "mysql", "userName"),
		password: config.ReadConfigByKey("./init.ini", "mysql", "password"),
		database: config.ReadConfigByKey("./init.ini", "mysql", "database"),
		url:      "",
	}
	m.url = fmt.Sprintf("%s:%s@%s(%s)/%s", m.userName, m.password, m.network, m.host, m.database)
	gloableMysqlHelper = m
	// db_mysql, err = sql.Open("mysql", conf.MysqlConfig.Url)
	// db_mysql.SetMaxOpenConns(1000)
	// db_mysql.SetMaxIdleConns(10)
	// db_mysql.SetConnMaxLifetime(time.Second * 10)
}

func (m *mysqlHelper) Test() {
	db, err := sql.Open("mysql", (*m).url)
	fmt.Println(err)
	defer db.Close()
	rows, err := db.Query("select id,orderNum from kh_orders limit 2")
	fmt.Println(err)
	for rows.Next() {
		var id string
		var orderNum string
		if err := rows.Scan(&id, &orderNum); err != nil {
			fmt.Println("error in rows")
			fmt.Println(err)
		}
		fmt.Println(id, orderNum)
	}
	rows.Close()
}
