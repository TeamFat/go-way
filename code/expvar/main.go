package main

import (
	"expvar"
	"fmt"
	"net/http"
)

var visits = expvar.NewInt("visits")

var stats = expvar.NewMap("tcp")
var requests, requestsFailed expvar.Int

func init() {
	stats.Set("requests", &requests)
	stats.Set("requests_failed", &requestsFailed)
}

func handler(w http.ResponseWriter, r *http.Request) {

	visits.Add(1)

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":1818", nil)
}

//https://orangetux.nl/post/expvar_in_action/
