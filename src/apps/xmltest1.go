package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"lib/publib/github.com/wonderivan/logger"
	"log"
	"net/http"
	"os"
	"testing/iotest"
)

//xml文件中节点
type StrResources struct {
	XMLName xml.Name `xml:"azkaban-users"` //标签名称
	Users   User   `xml:"user"`          //标签名称
}

type User struct {
	XMLName   xml.Name `xml:"user"`          //标签名称
	UserName  string   `xml:"username,attr"` //属性
	Roles     string   `xml:"roles,attr"`
	Password  string   `xml:"password,attr"`
	Groups    string   `xml:"groups,attr"`
	InnerText string   `xml:",innerxml"` //内部的文本
}

func encodeXml() {
	//获取xml文件位置
	file := fmt.Sprintf("%s/file/xmltest.xml",os.Getenv("GOPATH"))
	logger.Debug("file=%s",file)
	var r io.Reader
	decoder := xml.NewDecoder(iotest.NewReadLogger(file,r))
	var result StrResources
	err :=decoder.Decode(result)
	if err != nil {
		log.Println("Decode Fail : ",err)
		os.Exit(1)
	}
	fmt.Println("result.ResourceString:",result.Users)
	log.Println("UserName:",result.XMLName)
	log.Println("UserName:",result.Users.UserName)
	log.Println("Roles:",result.Users.Roles)
	log.Println("Password:",result.Users.Password)
	log.Println("Groups:",result.Users.Groups)
	log.Println("InnerText:",result.Users.InnerText)

	http.RoundTripper()
	http.CookieJar()
	http.Transport{}

}

func main() {

	encodeXml()
}
