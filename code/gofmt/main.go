package main

import (
	"go/format"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("gofmt", func(code string) string {
		gofmted, _ := format.Source([]byte(code))
		return string(gofmted)
	})
	//go get -u github.com/gopherjs/gopherjs
	// gopherjs build --minify
}
