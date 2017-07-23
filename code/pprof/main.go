package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"regexp"
)

func handler(wr http.ResponseWriter, r *http.Request) {
	var pattern = regexp.MustCompile(`^(\w+)@didichuxing.com$`)
	account := r.URL.Path[1:]
	res := pattern.FindSubmatch([]byte(account))
	if len(res) > 1 {
		wr.Write(res[1])
	} else {
		wr.Write([]byte("None"))
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	//http: //xargin.com/pprof-he-huo-yan-tu/
}
