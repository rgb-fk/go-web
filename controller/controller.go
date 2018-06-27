package controller

import (
	"fmt"
	"net/http"

	"github.com/everywan/go-web/config"
	"github.com/everywan/go-web/service"
)

func StartWebServer() {
	defer func() {
		if err := recover(); err != nil {
			logger_custom.ERROR(fmt.Sprintf("启动socket网络服务(startWebServer) 发生错误: %+v", err))
		}
	}()

	// 绑定静态页面
	// http.Handle("/", http.FileServer(http.Dir(htmlDir)))

	http.HandleFunc("/", service.FileServer)

	// 微信
	http.HandleFunc(fmt.Sprintf("/%s/%s", wx, "handle"), service.WxHandle)

	// 显示内容目录
	http.HandleFunc(fmt.Sprintf("/%s/%s", version, "lsDir"), service.LsDir)
	http.HandleFunc(fmt.Sprintf("/%s/%s/", version, "getFile"), service.GetFile)
	http.HandleFunc(fmt.Sprintf("/%s/%s/", version, "saveFile"), service.SaveFile)
	http.HandleFunc(fmt.Sprintf("/%s/%s/", version, "delFile"), service.DelFile)

	logger_custom.Info("服务启动成功")

	// 设置监听
	listenPort := config.ReadConfigByKey("./init.ini", "Net", "listenPort")
	err := http.ListenAndServe(":"+listenPort, nil)
	if err != nil {
		logger_custom.ERROR(fmt.Sprintf("启动socket网络服务(startWebServer) 发生错误(ListenAndServe方法): %+v", err))
	}
}
