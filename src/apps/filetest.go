/*
	1.测试命令行参数
	2.测试文件操作
*/
package main

import (
	"flag"
	"fmt"
	"lib/applib/myfile"
	"lib/publib/github.com/wonderivan/logger"
	"os"
	)

var infile *string = flag.String("i", "", "file contain values to sort")
var outfile *string = flag.String("o", "", "file to store sorted values")

func main() {
	var myinfile, myoutfile myfile.MyFile

	flag.Parse()
	if infile != nil && outfile != nil && len(*infile) > 0 && len(*outfile) > 0 {
		logger.Debug("infile=%s outfile=%s", *infile, *outfile)
	} else {
		logger.Error("Usage: sort -i <infile> -o <outfile>")
		return
	}

	dir := fmt.Sprintf("%s/file", os.Getenv("GOPATH"))
	myinfile.MyfileInit(dir, *infile)
	myinfile.ReadFileValues()
	logger.Debug("myinfile.Values=%v", myinfile.Values)

	myinfile.SortChan()
	logger.Debug("myinfile.Values=%v", myinfile.Values)

	myoutfile.MyfileInit(dir, *outfile)
	myoutfile.Values = myinfile.Values
	logger.Debug("myoutfile.Values=%v", myoutfile.Values)

	myoutfile.WriteFileValue()

	return
}
