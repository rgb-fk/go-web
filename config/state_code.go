package config

type Code struct {
	Code int
	Msg  string
}

var FAILD *Code = &Code{
	Code: 0,
	Msg:  "执行失败",
}

var SUCCEED *Code = &Code{
	Code: 1,
	Msg:  "成功",
}

var READFILEFAIL *Code = &Code{
	Code: 101,
	Msg:  "读取文件错误",
}

var FILENOTFOUND *Code = &Code{
	Code: 102,
	Msg:  "文件不存在",
}

var PASSWDNOTPAIR *Code = &Code{
	Code: 201,
	Msg:  "密码不匹配",
}
