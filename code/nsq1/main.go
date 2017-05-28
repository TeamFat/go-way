package main

import (
	"flag"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	go startConsumer()
	startProducer()
}

var url string
var url1 string

func init() {
	flag.StringVar(&url, "url", "127.0.0.1:4150", "nsqd")         //tcp
	flag.StringVar(&url1, "url1", "127.0.0.1:4161", "nsqlookupd") //http

	flag.Parse()
}

// 生产者
func startProducer() {
	cfg := nsq.NewConfig()
	producer, err := nsq.NewProducer(url, cfg)
	if err != nil {
		log.Fatal(err)
	}
	// 发布消息
	for {
		if err := producer.Publish("test", []byte("test message")); err != nil {
			log.Fatal("publish error: " + err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}

// 消费者
func startConsumer() {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "sensor01", cfg)
	if err != nil {
		log.Fatal(err)
	}
	// 设置消息处理函数
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println(string(message.Body))
		return nil
	}))
	// nsqlookupd
	//[]string
	if err := consumer.ConnectToNSQLookupds([]string{url1}); err != nil {
		log.Fatal(err)
	}
	<-consumer.StopChan
	//更多
	//http://tleyden.github.io/blog/2014/11/12/an-example-of-using-nsq-from-go/
	//https://www.yryz.net/post/go-nsq-usage.html
}
