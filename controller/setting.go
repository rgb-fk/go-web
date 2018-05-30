package controller

import (
	"github.com/everywan/go-web-demo/config"
	"github.com/everywan/go-web-demo/log"
)

var version string
var logger_custom log.Log
var htmlDir string

const wx = "wx"

func init() {
	version = config.ReadConfigByKey("./init.ini", "info", "version")
	htmlDir = config.ReadConfigByKey("./init.ini", "Work", "htmlDir")
	logger_custom.Init()
}
