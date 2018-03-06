package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// conenct zookeeper cluster
func connect_zk() (*zk.Conn, error) {
	c, _, err := zk.Connect([]string{"10.187.194.161:2181", "10.187.194.164:2181", "10.187.194.162:2181"}, time.Second*1)
	// c, session, err := zk.Connect([]string{"127.0.0.1:2182"}, time.Second*1)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func delete_node(c *zk.Conn, nodePath string) {
	c.Delete(nodePath, -1)
}

func create_node(c *zk.Conn, nodePath string) error {
	if _, err := c.Create(nodePath, []byte{1, 2, 3, 4}, 0, zk.WorldACL(zk.PermAll)); err != nil {
		fmt.Printf("Create returned error: %+v\n", err)
		return err
	}
	return nil
}

func get_node(c *zk.Conn, path string) {
	data, stat, err := c.Get(path)
	if err != nil {
		fmt.Printf("Get returned error: %+v\n", err)
	} else {
		fmt.Printf("Get node, data: %+v, state: %+v\n", data, stat)
	}
}

func watch_node(c *zk.Conn) {
	var wg sync.WaitGroup
	path := "/chenjie-2"

	delete_node(c, path)

	childChs, _, childCh, _ := c.ChildrenW("/")
	fmt.Println(childChs)
	/*
		if err != nil {
			fmt.Printf("Watch error: %v\n", err)
			return
		}
	*/

	wg.Add(1)
	go Watcher(wg, childCh)

	create_node(c, path)

	/*
		if path, err = c.Create("/gozk-test-2", []byte{1, 2, 3, 4}, 0, zk.WorldACL(zk.PermAll)); err != nil {
			fmt.Printf("Creat node: %v error: %v\n", "/gozk-test-2", err)
			return
		} else if path != "/gozk-test-2" {
			fmt.Printf("Create returned different path '%s' != '%s'", path, "/gozk-test-2")
			return
		}
	*/

	// add watcher for new added node
	_, _, addCh, _ := c.ChildrenW("/chenjie-2")
	wg.Add(1)
	go Watcher(wg, addCh)

	delete_node(c, path)

	wg.Wait()
}

func main() {
	path := "/chenjie"
	// connect zk-server
	c, err := connect_zk()
	if err != nil {
		return
	}

	// create zk-node, delete it if already exist
	delete_node(c, path)
	err = create_node(c, path)
	if err != nil {
		return
	}

	get_node(c, path)

	watch_node(c)
}

func Watcher(wg sync.WaitGroup, childCh <-chan zk.Event) {
	defer wg.Done()
	select {
	case ev := <-childCh:
		if ev.Err != nil {
			fmt.Printf("Child watcher error: %+v\n", ev.Err)
			return
		}
		fmt.Printf("Receive event: %+v\n", ev)
	case _ = <-time.After(time.Second * 2):
		fmt.Printf("Child Watcher timeout")
		return
	}
}
