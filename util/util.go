package util

import (
	"io/ioutil"
	"net/http"
	"os"
)

const GFM_URL = "https://api.github.com/markdown/raw"

func RenderGFM(path string) (string, error) {
	mdFile, err := os.Open(path)
	defer mdFile.Close()
	if err == nil {
		// 请求 gfm api ,渲染md文件
		resp, _ := http.Post(GFM_URL, "text/plain", mdFile)
		if resp.StatusCode == 200 {
			content, _ := ioutil.ReadAll(resp.Body)
			return string(content), nil
		}
	}
	return "", err
}

func VerifyPasswd(passwd string, _passwd string) bool {
	return passwd == _passwd
}
