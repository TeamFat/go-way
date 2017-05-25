package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("awaiting signal")
	<-sigs
	fmt.Println("exiting")
}
