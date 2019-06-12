package main

import "lib/publib/github.com/wonderivan/logger"

func main() {
	s := make([]int, 4, 4)
	logger.Debug("s.len=%v s.cap=%v", len(s), cap(s))
	s1 := append(s, 1)
	logger.Debug("s.len=%v s1.len=%v", len(s), len(s1))
	logger.Debug("s.cap=%v s1.cap=%v", cap(s), cap(s1))
}
