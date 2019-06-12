package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"lib/publib/github.com/wonderivan/logger"
	"log"
	"os"
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
	Age 		int		`xml:"age,attr"`
	Salary		float32	`xml:"salary"`
}

func encodeXml() {
/*	data := ` <?xml version="1.0" encoding="UTF-8"?>
<azkaban-users>
  <user groups="azkaban" password="azkabanqweqe" roles="admin" username="azkaban">The demo</user>
</azkaban-users> `*/
	//获取xml文件位置
	file := fmt.Sprintf("%s/file/xmltest.xml",os.Getenv("GOPATH"))
	logger.Debug("file=%s",file)
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("读文件位置错误信息：",err)
	}
	log.Printf("read contents: %s",content)

	//将文件转换为对象
	var result StrResources
	err = xml.Unmarshal(content, &result)
	if err != nil {
		fmt.Println("读文件内容错误信息：",err)
	}


	fmt.Println("result:",result)
	fmt.Println("result.ResourceString:",result.Users)
		log.Println("UserName:",result.XMLName)
		log.Println("UserName:",result.Users.UserName)
		log.Println("Roles:",result.Users.Roles)
		log.Println("Password:",result.Users.Password)
		log.Println("Groups:",result.Users.Groups)
		log.Println("InnerText:",result.Users.InnerText)

	fmt.Println("读取完成。。。。")
	result.Users.Age += 1
	result.Users.Salary += 5000

	out,err := xml.MarshalIndent(result,"","  ")
	tt := append([]byte(xml.Header),out...)
	err = ioutil.WriteFile(file,tt,755)
	//err = ioutil.WriteFile(file,out,755)
	if err != nil {
		log.Println("write file faile : ",err)
	}
}

func main() {

	encodeXml()
}
