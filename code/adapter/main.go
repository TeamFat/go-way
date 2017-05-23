package main

import "log"

type LogIN interface {
	WARN(string)
	INFO(string)
	//其他省略
}

/*
假设你写了一个框架,框架里需要打日志，具体是个什么样的日志由使用者决定
*/
type Frame struct {
	name string
	log  LogIN
}

func (f *Frame) Start() {
	f.log.INFO("frame start...")
}

//但是问题来了，使用者用了一个这样的日志库

type RealLog struct {
}

func (l *RealLog) warn(s string) {
	log.Print(s)
}

func (l *RealLog) info(s string) {
	log.Print(s)
}

//可以看到，真正使用的日志与框架定义的日志接口不一样，如此我们就需要适配

type AdaptLog struct {
	rellog *RealLog
}

func (alog *AdaptLog) WARN(s string) {
	alog.rellog.warn(s)
}

func (alog *AdaptLog) INFO(s string) {
	alog.rellog.info(s)
}

//有了这个适配器，我们就可以很好的在框架里使用RealLog了

func main() {
	adaptlog := AdaptLog{&RealLog{}}
	frame := Frame{name: "frame", log: &adaptlog}

	frame.Start()
}
