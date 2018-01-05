package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	curDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("Read path error: ", err.Error())
		return
	}
	fileDir := path.Join(curDir, "file-store")
	fmt.Println(fileDir)
	_ = os.Mkdir(fileDir, os.ModePerm) //0777
}
