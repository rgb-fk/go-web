package service

import (
	"github.com/everywan/go-web-demo/config"
	"github.com/everywan/go-web-demo/log"
)

// 微信使用token
var token string

// 网站操作密码
var passwd string

var contentDir string
var todoFile string
var logger log.Log
var htmlDir string

const wxReturn = "success"

func init() {
	token = config.ReadConfigByKey("./init.ini", "Work", "token")
	passwd = config.ReadConfigByKey("./init.ini", "Work", "passwd")
	contentDir = config.ReadConfigByKey("./init.ini", "Work", "contentDir")
	todoFile = config.ReadConfigByKey("./init.ini", "Work", "todoFile")
	htmlDir = config.ReadConfigByKey("./init.ini", "Work", "htmlDir")
	logger.Init()
}
