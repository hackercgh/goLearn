package hello

import (
	"lib/publib/github.com/wonderivan/logger"
)

func Hello() {

	//currentTimeData := time.Date(time.Now().Year(),time.Now().Month(),time.Now().Day(),time.Now().Hour(),time.Now().Minute(),time.Now().Second(),time.Now().Nanosecond(),time.Local).String()
	//log.Println( "hell-test: ",currentTimeData , " 马上放五一了，嗨起来！！！")
	//panic("测试panic")
	logger.Warn("马上放五一了，嗨起来！！！")
}

