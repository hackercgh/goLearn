package myfile

import (
	"bufio"
	"fmt"
	"io"
	"lib/publib/github.com/wonderivan/logger"
	"os"
	"strconv"
)

type MyFile struct {
	FileDir		string
	FileName	string
	Values 		[]int
}
func (myfile *MyFile) MyfileInit(dir string,filename string){
	myfile.FileDir 	= 	dir
	myfile.FileName	=	filename
	myfile.Values	=	make([]int,0)
}
func (myfile *MyFile)ReadFileValues() ( err error) {
	dirFile := fmt.Sprintf("%s/%s",myfile.FileDir,myfile.FileName)
	file, err := os.OpenFile(dirFile, os.O_RDONLY, 755)
	if err != nil {
		logger.Error("os.OpenFile faile,err=[%v]", err)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				logger.Error("br.ReadLine faile,err=[%s]", err1)
				err = err1
				return
			}
			break
		}

		if isPrefix {
			logger.Error("the line was too long for the buffer and isPrefix is seted ")
		}

		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			logger.Error("trconv.Atoi(str) faile, err=[%v", err1)
			return
		}
		myfile.Values = append(myfile.Values, value)
	}

	return
}

func (myfile *MyFile)SortChan() (err error) {
	iLen := len(myfile.Values)
	for i:=0;i<iLen-1;i++ {
		for j:=0;j<iLen-i-1;j++{
			if myfile.Values[j] > myfile.Values[j+1] {
				myfile.Values[j],myfile.Values[j+1] = myfile.Values[j+1],myfile.Values[j]
			}
		}
	}
	return nil
}

func (myfile *MyFile)WriteFileValue() (err error) {
	dirFile := fmt.Sprintf("%s/%s",myfile.FileDir,myfile.FileName)
	//file,err := os.Create(dirFile)
	//if file,err := os.OpenFile(dirFile,os.O_CREATE | os.O_TRUNC|os.O_WRONLY,755); err != nil{
	 file,err := os.OpenFile(dirFile,os.O_CREATE | os.O_TRUNC|os.O_WRONLY,755)
	if err != nil {
		logger.Error("err=[%v]",err)
		return
	} else {
		logger.Debug("OpenFile ok")
	}

	defer  file.Close()

	for _,value := range  myfile.Values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
}
