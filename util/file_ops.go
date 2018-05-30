package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func GetFileHelper(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	return string(data), err
}

func SaveFileHelper(path string, content string) (string, error) {
	// 判断路径是否存在
	if !FileExist(path) {
		path = "/home/wzs/website/doc/tmp/" + fmt.Sprint(time.Now().Unix()) + ".md"
	}
	err := ioutil.WriteFile(path, []byte(content), 0644)
	return path[21:], err
}

func DelFileHelper(path string) (string, error) {
	err := os.Remove(path)
	return path[21:], err
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}
