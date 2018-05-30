package service

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/everywan/go-web-demo/entity"
	"github.com/everywan/go-web-demo/util"
)

func WxHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result := wxReturn
	defer func() {
		if err := recover(); err != nil {
			logger.ERROR(fmt.Sprintf("执行 wxAuth 方法出错, 错误是: %+v", err))
		}
		logger.Debug1("发送消息; result:" + result)
		fmt.Fprintf(w, "%s", result)
	}()
	r.ParseForm()

	signature := r.FormValue("signature")
	timestamp := r.FormValue("timestamp")
	nonce := r.FormValue("nonce")
	echostr := r.FormValue("echostr")
	// 验证签名
	if util.MakeSignature(token, timestamp, nonce) != signature {
		result = "sign err"
		return
	}

	// POST表示非认证消息
	if r.Method == "POST" {
		xmlBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.ERROR("获取post数据失败, err: " + err.Error())
			return
		}
		logger.Debug1(fmt.Sprintf("收到微信消息; xmlBody: %s", xmlBody))
		reMsg, _ := wxReceive(xmlBody)
		result = wxReply(reMsg)
		return
	} else {
		logger.Debug1(fmt.Sprintf("收到微信认证消息; signature:%s, timestamp:%s, nonce:%s, echostr:%s", signature, timestamp, nonce, echostr))
		result = echostr
		return
	}
}

func wxReceive(xmlBody []byte) (*entity.WxReceiveMsg, error) {
	data := &entity.WxReceiveMsg{}
	err := xml.Unmarshal(xmlBody, &data)
	if err != nil {
		logger.ERROR("解析xml失败, err: " + err.Error())
		return nil, err
	}
	return data, nil
}

func wxReply(msg *entity.WxReceiveMsg) string {
	result := wxReturn
	// 通用回复类型
	reply := &entity.WxReplyMsg{
		CreateTime:   time.Now().Unix(),
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		MsgType:      msg.MsgType,
	}
	switch msg.Content {
	case "todo":
		{
			reply.Content, _ = util.GetFileHelper("/home/wzs/website/doc/todo.md")
		}
	default:
		{
			reply.Content = "还没想好自动回复什么比较好"
		}
	}
	if _xmlBody, err := xml.Marshal(reply); err == nil {
		result = string(_xmlBody)
	}
	return result
}
