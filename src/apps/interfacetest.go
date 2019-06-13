/*
	1.学习interface使用
*/
package main

import (
	"lib/publib/github.com/wonderivan/logger"
)

/*实现接口play的MP3类型*/
type MP3Playper struct {
	musicNmae string
	singer    string
	album     string
}

func (mp3 *MP3Playper) play() {
	logger.Debug("this is MP3Player   musicName=%s singer=%s album=%s", mp3.musicNmae, mp3.singer, mp3.album)
}

/*实现接口play的WAV类型*/
type WAVPlayper struct {
	musicNmae string
	singer    string
	album     string
}

	func (wav *WAVPlayper) play() {
	logger.Debug("this is WAVPlayer   musicName=%s singer=%s album=%s", wav.musicNmae, wav.singer, wav.album)
}

type Player interface {
	play()
}

func myplay(p Player) {
	p.play()
}
func main() {
	mp3 := &MP3Playper{"成都\t", "赵雷\t\t", "流浪\t"}
	wav := &WAVPlayper{"怀念青春\t", "旭日阳刚\t", "单曲\t"}
	myplay(mp3)
	myplay(wav)
}
