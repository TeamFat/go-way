package main

/**
客户端doc地址：github.com/samuel/go-zookeeper/zk
**/
import (
	"fmt"
	"time"

	zk "github.com/samuel/go-zookeeper/zk"
)

/**
 * 获取一个zk连接
 * @return {[type]}
 */
func getConnect(zkList []string) (conn *zk.Conn) {
	conn, _, err := zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

/**
 * 测试连接
 * @return
 */
func test1() {
	zkList := []string{"10.187.194.161:2181", "10.187.194.164:2181", "10.187.194.162:2181"}
	conn := getConnect(zkList)

	defer conn.Close()
	conn.Create("/chenjie", nil, 0, zk.WorldACL(zk.PermAll))

	time.Sleep(20 * time.Second)
}

/**
 * 测试临时节点
 * @return {[type]}
 */
func test2() {
	zkList := []string{"10.187.194.161:2181", "10.187.194.164:2181", "10.187.194.162:2181"}
	conn := getConnect(zkList)

	defer conn.Close()
	conn.Create("/chenjie/chenjie1", []byte("hello,this is a zk go test demo!!!"), 0, zk.WorldACL(zk.PermAll))

	time.Sleep(20 * time.Second)
}

/**
 * 获取所有节点
 */
func test3() {
	zkList := []string{"10.187.194.161:2181", "10.187.194.164:2181", "10.187.194.162:2181"}
	conn := getConnect(zkList)

	defer conn.Close()

	children, _, err := conn.Children("/chenjie")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v \n", children)
}

func main() {
	test1()
	test2()
	test3()
}
