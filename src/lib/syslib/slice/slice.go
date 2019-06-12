package slice

import (
	"lib/publib/github.com/wonderivan/logger"
)
func SliceTest(){
	s := make([] int,5)
	sliceInit(s)
	logger.Debug("slice ",s)
	sliceExchang(s)
	logger.Debug("slice ",s)
}

func sliceInit(s []int )  {
	for i:= 0;i < len(s);i++{
		s[i] = i
	}
}

func sliceExchang(s []int)   {
	iLen := len(s)
	for i := 0;i< iLen/2;i++ {
		s[i],s[iLen-i - 1] = s[iLen-i -1],s[i]
	}
}
