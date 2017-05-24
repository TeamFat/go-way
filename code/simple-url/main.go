package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resq, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resq.Body)
		resq.Body.Close()
		if err != nil {
			fmt.Println(url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
