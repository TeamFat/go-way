package main

import (
	"fmt"

	"github.com/bitly/go-simplejson"
)

func main() {
	js, err := simplejson.NewJson([]byte(`{
		"test": {
			"array": ["_", 3],
			"int": 10,
			"float": 5.150,
			"bignum": 9223372036854775807,
			"string": "simplejson",
			"bool": true
		}
	}`))
	if err != nil {

	}
	arr, _ := js.Get("test").Get("array").Array()
	//i, _ := js.Get("test").Get("int").Int()
	//ms := js.Get("test").Get("string").MustString()
	fmt.Println(arr[0])
	fmt.Println(arr[1])

	//输出
	//_
	//3
}
