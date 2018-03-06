package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

var hosts = []string{"10.187.194.161:2181", "10.187.194.164:2181", "10.187.194.162:2181"}

var path1 = "/chenjie2"

var flags int32 = zk.FlagEphemeral
var data1 = []byte("hello,this is a zk go test demo!!!")
var acls = zk.WorldACL(zk.PermAll)

func main() {
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	_, _, ech, err := conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}

	create(conn, path1, data1)

	go watchCreataNode(ech)
	// 	部分监听：
	// 1.调用conn.ExistsW(path) 或GetW(path)为对应节点设置监听，该监听只生效一次
	// 2.开启一个协程处理chanel中传来的event事件
	// （注意：watchCreataNode一定要放在一个协程中，不能直接在main中调用，不然会阻塞main）
	//1.如果即设置了全局监听有设置了部分监听，那么最终是都会触发的，并且全局监听在先执行
	//2.如果设置了监听子节点，那么事件的触发是先子节点后父节点
}

func create(conn *zk.Conn, path string, data []byte) {
	_, err_create := conn.Create(path, data, flags, acls)
	if err_create != nil {
		fmt.Println(err_create)
		return
	}

}

func watchCreataNode(ech <-chan zk.Event) {
	event := <-ech
	fmt.Println("*******************")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("-------------------")
}
