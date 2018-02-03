//如下：
package main

import (
    "fmt"
    "github.com/samuel/go-zookeeper/zk"
    "time"
)

var hosts = []string{"115.159.215.179:2181"}


func main() {
	//http://izualzhy.cn/c/cpp/2016/09/24/zkcli-introduction
	//https://birdben.github.io/2016/09/01/Docker/Docker%E5%AE%9E%E6%88%98%EF%BC%88%E5%8D%81%E4%B8%83%EF%BC%89Docker%E5%AE%89%E8%A3%85Zookeeper%E7%8E%AF%E5%A2%83/
	conn, _, err := zk.Connect(hosts, time.Second*5)
    if err != nil {
        fmt.Println(err)
        return
    }
	defer conn.Close()
	// var path="/test"
	// var data=[]byte("hello zk")
	// var flags=int32(0)
	// //flags有4种取值：
	// //0:永久，除非手动删除
	// //zk.FlagEphemeral = 1:短暂，session断开则改节点也被删除
	// //zk.FlagSequence  = 2:会自动在节点后面添加序号
	// //3:Ephemeral和Sequence，即，短暂且自动添加序号
	// var acls=zk.WorldACL(zk.PermAll)//控制访问权限模式

	// p,err_create:=conn.Create(path,data,flags,acls)
	// if err_create != nil {
    // 	fmt.Println(err_create)
    // 	return
	// }
	// fmt.Println("create:",p)
	children, stat, ch, err := conn.ChildrenW("/test")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v %+v\n", children, stat)
	e := <-ch
	fmt.Printf("%+v\n", e)
}
