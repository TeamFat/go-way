package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://baidu.com")
	//bl := strings.HasPrefix("http://www.baidu.com", "http:")
	//fmt.Println(bl)
	//io.Copy(os.Stdout, resp.Body)
	//fmt.Println(resp.Status)
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(b))
	fmt.Printf("%s", b)

}
