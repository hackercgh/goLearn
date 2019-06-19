package main

import (
	"fmt"
	"io/ioutil"
	"lib/publib/github.com/wonderivan/logger"
	"log"
	"net/http"
	"os"
)

func main() {

	logger.Debug("8 << 1 = %v",8 << 1)
	logger.Debug("8 >> 1 = %v",8 >> 1)
	ExampleFileServer()

}

func ExampleGet() {
	res, err := http.Get("https://www.jianshu.com")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func ExampleFileServer() {
	// Simple static webserver:
	dir := fmt.Sprintf("%s/file",os.Getenv("GOPATH"))
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(dir))))
}
