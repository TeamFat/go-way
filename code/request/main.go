package main

import (
	"fmt"
	"net/http"

	"github.com/mozillazg/request"
)

func main() {
	c := new(http.Client)
	req := request.NewRequest(c)
	req.Data = map[string]string{
		"key": "value",
		"a":   "123",
	}
	resp, _ := req.Post("http://httpbin.org/post")
	j, _ := resp.Json()
	defer resp.Body.Close() // Don't forget close the response body
	fmt.Println(j)
}
