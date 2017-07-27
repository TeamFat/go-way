package main

import (
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

func child() {
	cmd := exec.Command("sleep", "600")
	//start := time.Now()
	time.AfterFunc(10*time.Second, func() { cmd.Process.Kill() })
	cmd.Run()
}

// func grand_child() {
// 	cmd := exec.Command("/bin/sh", "-c", "sleep 1000")
// 	time.AfterFunc(10*time.Second, func() { cmd.Process.Kill() })
// 	cmd.Run()
// }

func grand_child() {
	cmd := exec.Command("/bin/sh", "-c", "sleep 1000")
	// Go会将PGID设置成与PID相同的值
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	start := time.Now()
	time.AfterFunc(10*time.Second, func() { syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL) })
	err := cmd.Run()
	fmt.Printf("pid=%d duration=%s err=%s\n", cmd.Process.Pid, time.Since(start), err)
}

func main() {
	child()
	grand_child()
	//http://www.jianshu.com/p/1f3ec2f00b03
	//http://colobu.com/2015/10/09/Linux-Signals/
}
