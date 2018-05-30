package log

import (
	"fmt"
	"time"

	"github.com/everywan/go-web-demo/config"
)

// 时间格式
const timeLayout = "2006-01-02 03:04:05"

type Log struct {
	stdout bool
	file   string
	level  string
}

// 没有采用 init() 单例 是为了后续可以自定义不同的log实现差异化输出
func (logger *Log) Init() {
	if config.ReadConfigByKey("./init.ini", "Log", "stdout") == "true" {
		logger.stdout = true
		return
	} else {
		logger.stdout = false
	}
	logger.file = config.ReadConfigByKey("./init.ini", "Log", "file")
	logger.level = config.ReadConfigByKey("./init.ini", "Log", "level")
}

func (logger *Log) Info(msg interface{}) {
	fmt.Printf("Info: %s: %v \n", time.Now().Format(timeLayout), msg)
}

func (logger *Log) ERROR(msg interface{}) {
	fmt.Printf("ERROR: %s: %v \n", time.Now().Format(timeLayout), msg)
}

func (logger *Log) WARNING(msg interface{}) {
	fmt.Printf("WARNING: %s: %v \n", time.Now().Format(timeLayout), msg)
}

func (logger *Log) Debug1(msg interface{}) {
	fmt.Printf("Debug1: %s: %v \n", time.Now().Format(timeLayout), msg)
}
func (logger *Log) Debug2(msg interface{}) {
	fmt.Printf("Debug2: %s: %v \n", time.Now().Format(timeLayout), msg)
}
func (logger *Log) Debug3(msg interface{}) {
	fmt.Printf("Debug3: %s: %v \n", time.Now().Format(timeLayout), msg)
}

func (logger *Log) Log(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %s; err: %s \n", time.Now().Format(timeLayout), msg, err.Error())
		panic(err)
	}
}

func (logger *Log) Test(msg interface{}) {
	fmt.Println("----------------------------TEST_BEGIN----------------------------")
	fmt.Printf("Test: %s: %v \n", time.Now().Format(timeLayout), msg)
	fmt.Println("----------------------------TEST_END---------------------------")
}

func (logger *Log) MakeError(msg interface{}) error {
	return fmt.Errorf("%s: %v", time.Now().Format(timeLayout), msg)
}

func MakeError(msg interface{}) error {
	return fmt.Errorf("%s: %v", time.Now().Format(timeLayout), msg)
}
