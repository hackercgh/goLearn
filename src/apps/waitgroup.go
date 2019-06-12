package main

import (
	"lib/publib/github.com/wonderivan/logger"
	"runtime"
	"sync"
)

const RuntimeNum = 5
func main() {
	logger.Warn("***************************************************************************************************")
	logger.Warn("1.学习使用服务组进行gorutine同步")
	logger.Warn("***************************************************************************************************")

	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := sync.WaitGroup{} //创建服务组
	for i := 0; i < RuntimeNum; i++ {
		wg.Add(1)	//Add adds delta to the WaitGroup counter
		go Go(&wg, i)
	}
	wg.Wait()	//Wait blocks until the WaitGroup counter is zero.
}

/*gorutine服务组计数器消费*/
func Go (wg *sync.WaitGroup,idx int) {
	//logger.Debug("传入idx = %d",idx)
	for i:= 0 ;i< 10;i++ {
		idx += i
	}
	logger.Debug("idx = %v",idx)
	wg.Done()	//Done decrements the WaitGroup counter by one.
}