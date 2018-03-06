package example

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

const (
	timeOut = 20
)

var hosts []string = []string{"10.187.194.161:2181", "10.187.194.164:2181", "10.187.194.162:2181"} // the zk server list

func GetConnect() (conn *zk.Conn, err error) {
	conn, _, err = zk.Connect(hosts, timeOut*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func RegistServer(conn *zk.Conn, host string) (err error) {
	_, err = conn.Create("/chenjie/"+host, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return
}

func GetServerList(conn *zk.Conn) (list []string, err error) {
	list, _, err = conn.Children("/chenjie")
	return
}
