package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

var hosts = []string{"10.187.194.161:2181", "10.187.194.164:2181", "10.187.194.162:2181"}

var path1 = "/chenjie1"

var flags int32 = zk.FlagEphemeral
var data1 = []byte("hello,this is a zk go test demo!!!")
var acls = zk.WorldACL(zk.PermAll)

func main() {
	// 	Java API中是通过Watcher实现的，在go-zookeeper中则是通过Event。道理都是一样的
	// 全局监听：
	// 1.调用zk.WithEventCallback(callback)设置回调
	// 2.调用conn.ExistsW(path) 或GetW(path)为对应节点设置监听，该监听只生效一次
	option := zk.WithEventCallback(callback)

	conn, _, err := zk.Connect(hosts, time.Second*5, option)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	_, _, _, err = conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}

	create(conn, path1, data1)

	time.Sleep(time.Second * 2)

	_, _, _, err = conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}

}
func callback(event zk.Event) {
	fmt.Println("*******************")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("-------------------")
}

func create(conn *zk.Conn, path string, data []byte) {
	//flags有4种取值：
	//0:永久，除非手动删除
	//zk.FlagEphemeral = 1:短暂，session断开则改节点也被删除
	//zk.FlagSequence  = 2:会自动在节点后面添加序号
	//3:Ephemeral和Sequence，即，短暂且自动添加序号
	//var acls=zk.WorldACL(zk.PermAll)//控制访问权限模式
	_, err_create := conn.Create(path, data, flags, acls)
	if err_create != nil {
		fmt.Println(err_create)
		return
	}

}
