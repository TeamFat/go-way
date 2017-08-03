package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.ListenAndServe("localhost:13001", nil)
}
