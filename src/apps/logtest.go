package main

import (
	"lib/applib/hello"
	"lib/publib/github.com/wonderivan/logger"
	"lib/syslib/slice"
)

//var p = `{
//	"TimeFormat":"2006-01-02 15:04:05", // 输出日志开头时间格式
//	"Console": {         // 控制台日志配置
//		"level": "TRAC",    // 控制台日志输出等级
//		"color": true       // 控制台日志颜色开关
//	},
//	"File": {                   // 文件日志配置
//		"filename": "app.log",  // 初始日志文件名
//		"level": "TRAC",        // 日志文件日志输出等级
//		"daily": true,          // 跨天后是否创建新日志文件，当append=true时有效
//		"maxlines": 1000000,    // 日志文件最大行数，当append=true时有效
//		"maxsize": 1,           // 日志文件最大大小，当append=true时有效
//		"maxdays": -1,          // 日志文件有效期
//		"append": true,         // 是否支持日志追加
//		"permit": "0660"        // 新创建的日志文件权限属性
//	}
//}`
var p = `{
	"Console": {
		"level": "DEBG",
		"color": true
	},
	"File": {  
		"filename": "app.log",
		"level": "EROR",
		"daily": true,
		"maxlines": 1000000,
		"maxsize": 256,
		"maxdays": -1,
		"append": true,
		"permit": "0660"
	}
}`

func main() {
	logger.SetLogger()
	defer func() {
		if r := recover(); r != nil {
			logger.Error("进程从panic中recover,r=%v", r)
		}
	}()
	logger.Error("this is Error")
	hello.Hello()
	logger.Error("--------------------------")
	slice.SliceTest()
	logger.Error("--------------------------")
	//logger.Register()
	//fmt.Println(os.Getpid())
}
