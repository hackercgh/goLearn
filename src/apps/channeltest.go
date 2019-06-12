/*
	1.学习goroutin 间的chan数据传输(进程间通讯可使用分布式系统的方式解决，比如使用socket、http通讯协议)
	2.学习select使用
*/
package main

import (
	"lib/publib/github.com/wonderivan/logger"
	"runtime"
	"sync"
	"time"
)

const TimeOut = 3

func main() {
	logger.Warn("***************************************************************************************************")
	logger.Warn("1.学习goroutin 间的chan数据传输(进程间通讯可使用分布式系统的方式解决，比如使用socket、http通讯协议)")
	logger.Warn("2.学习select使用")
	logger.Warn("***************************************************************************************************")
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}

	iChan := make(chan int, 1)
	sChan := make(chan rune, 1)

	//iChan 生产者
	go sendIchan(&wg, iChan)
	//sChan生产者
	go sendSchan(&wg, sChan)
	//iChan,sChan消费者
	go mainComm(&wg, iChan, sChan, TimeOut)
	time.Sleep((TimeOut + 1) * time.Second)

	//等待工作组结束后退出主进程
	wg.Wait()

	//	关闭channel
	close(iChan)
	close(sChan)
}

/*单向channel 只生产不消费*/
func sendIchan(wg *sync.WaitGroup, iChan chan<- int) {
	for i := 0; i < 10; i++ {
		wg.Add(1) //正确
		iChan <- i
		//wg.Add(1) // 在外层done时会有问题
		//time.Sleep(1*time.Second)
		//logger.Debug("add %v to iChan",i)
	}
	return
}

/*单向channel 只生产不消费*/
func sendSchan(wg *sync.WaitGroup, sChan chan<- rune) {
	for c := 'A'; c <= 'Z'; c++ {
		wg.Add(1)
		sChan <- c
		//time.Sleep(1*time.Second)
		//logger.Debug("add %c to sChan",c)
	}
	return
}

/*单向channel 只消费*/
func mainComm(wg *sync.WaitGroup, iChan <-chan int, sChan <-chan rune, iTimeOut time.Duration) {

	for {
		select {
		case v, ok := <-iChan:
			if !ok {
				logger.Error("iChan Read Error")
				break
			}
			logger.Debug("read %v From iChan", v)
			wg.Done()
		case v, ok := <-sChan:
			if !ok {
				logger.Error("sChan Read Error")
				break
			}
			logger.Debug("read %c From sChan", v)
			wg.Done()
		case <-time.After(iTimeOut * time.Second):
			logger.Error("%ds 超时未收到数据", iTimeOut)
			break
		}
	}
}
