package main

import (
	"log"
	"os"
)

var cwd string

func init() {
	//cwd, err := os.Getwd()
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed:%v", err)
	}
	log.Printf("Working directory =%s", cwd)
}
func main() {

}
