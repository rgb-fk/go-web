package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/everywan/go-web/config"
	"github.com/everywan/go-web/entity"
	"github.com/everywan/go-web/util"
)

func FileServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_url := r.URL.Path
	_len := len(_url)
	// if r.Host == "resume.xiagaoxiawan.com" {
	// 	http.Redirect(w, r, "https://everywan.github.io/resume/", http.StatusFound)
	// 	return
	// }

	// 首先根据 Host 做一遍划分
	_htmlDir := func() string {
		switch r.Host {
		case "www.xiagaoxiawan.com":
			return htmlDir
		case "todo.xiagaoxiawan.com":
			return htmlDir + "/play/TODO"
		case "drawgraph.xiagaoxiawan.com":
			return htmlDir + "/play/DrawGraph"
		case "resume.xiagaoxiawan.com":
			return htmlDir + "/play/resume"
		case "private.xiagaoxiawan.com":
			// 还没想好
			// return htmlDir + "/play/resume"
		}
		return htmlDir
	}()

	if _url == "/" {
		// 默认情况
		_url = _htmlDir + "/index.html"
	} else if _len > 4 && _url[_len-4:] == ".css" {
		_url = _htmlDir + _url
		w.Header().Set("Content-Type", "text/css")
	} else {
		_url = _htmlDir + _url
	}
	data, _ := util.GetFileHelper(_url)
	fmt.Fprintf(w, "%s", data)
}

func GetFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result := &entity.Result_web{}
	result.TransCode(config.FAILD)
	defer func() {
		if err := recover(); err != nil {
			logger.ERROR(fmt.Sprintf("执行web方法出错(execCMD), 错误是: %+v", err))
		}
		jsonResult, _ := json.Marshal(result)
		fmt.Fprintf(w, "%s", jsonResult)
	}()
	rePath := r.URL.Path[11:]
	if data, err := util.GetFileHelper(contentDir + rePath); err == nil {
		result.Data = data
		result.TransCode(config.SUCCEED)
	}
}

func SaveFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result := &entity.Result_web{}
	result.TransCode(config.FAILD)
	defer func() {
		if err := recover(); err != nil {
			logger.ERROR(fmt.Sprintf("执行web方法出错(execCMD), 错误是: %+v", err))
		}
		jsonResult, _ := json.Marshal(result)
		fmt.Fprintf(w, "%s", jsonResult)
	}()
	_passwd := r.FormValue("passwd")
	if !util.VerifyPasswd(passwd, _passwd) {
		result.TransCode(config.PASSWDNOTPAIR)
		return
	}
	rePath := r.URL.Path[12:]
	content := r.PostFormValue("content")
	path, err := util.SaveFileHelper(contentDir+rePath, content)
	if err != nil {
		panic(err)
	}
	result.Data = path
	result.TransCode(config.SUCCEED)
}
func DelFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result := &entity.Result_web{}
	result.TransCode(config.FAILD)
	defer func() {
		if err := recover(); err != nil {
			logger.ERROR(fmt.Sprintf("执行web方法出错(execCMD), 错误是: %+v", err))
		}
		jsonResult, _ := json.Marshal(result)
		fmt.Fprintf(w, "%s", jsonResult)
	}()
	_passwd := r.FormValue("passwd")
	if !util.VerifyPasswd(passwd, _passwd) {
		result.TransCode(config.PASSWDNOTPAIR)
		return
	}
	rePath := r.URL.Path[11:]
	path, err := util.DelFileHelper(contentDir + rePath)
	if err != nil {
		panic(err)
	}
	result.Data = path
	result.TransCode(config.SUCCEED)
}

func LsDir(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result := &entity.Result_web{}
	result.TransCode(config.FAILD)
	defer func() {
		if err := recover(); err != nil {
			logger.ERROR(fmt.Sprintf("执行web方法出错(execCMD), 错误是: %+v", err))
		}
		jsonResult, _ := json.Marshal(result)
		fmt.Fprintf(w, "%s", jsonResult)
	}()
	result.Data = LsDirHelper(contentDir)
	result.TransCode(config.SUCCEED)
}
func LsDirHelper(dir string) []map[string]interface{} {
	relDir := strings.Replace(dir, contentDir, "", -1)
	dirs := strings.Split(relDir, "/")
	baseProject := ""
	relPathBaseProject := ""
	if len(dirs) > 1 {
		baseProject = "/" + dirs[1]
		relPathBaseProject = strings.Replace(relDir, "/"+dirs[1], "", 1)
	}
	var data []map[string]interface{}
	infos, _ := ioutil.ReadDir(dir)
	// 两次循环, 区分开目录和文件, 目录在前,文件在后
	for _, info := range infos {
		if info.IsDir() {
			fileName := info.Name()
			if fileName[0] == '.' {
				continue
			}
			var node = make(map[string]interface{}, 2)
			node["text"] = fileName
			node["nodes"] = LsDirHelper(dir + "/" + info.Name())
			data = append(data, node)
		}
	}
	for _, info := range infos {
		if info.IsDir() {
			continue
		}
		fileName := info.Name()
		if fileName[0] == '.' {
			continue
		}
		var node = make(map[string]interface{}, 2)
		node["text"] = fileName
		node["relPath"] = baseProject
		node["href"] = relPathBaseProject + "/" + fileName
		data = append(data, node)
	}
	return data
}
