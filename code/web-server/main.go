package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	// http://localhost:8000/count99?ff=23
	// URL.Path = "/count99"
	// GET /count99?ff=23 HTTP/1.1
	// Header["Connection"] = ["keep-alive"]
	// Header["Accept-Encoding"] = ["gzip, deflate, sdch, br"]
	// Header["Accept-Language"] = ["zh-CN,zh;q=0.8,en;q=0.6,zh-TW;q=0.4,es;q=0.2,ko;q=0.2"]
	// Header["Cookie"] = ["Phpstorm-9cf989ed=cdef96be-21e7-4622-b3c1-bd834205a2e0; Hm_lvt_1ad0a5c62a113bb874be8bb514b0a70d=1490672835"]
	// Header["Socketlog"] = ["SocketLog(tabid=457&client_id=)"]
	// Header["Upgrade-Insecure-Requests"] = ["1"]
	// Header["User-Agent"] = ["Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"]
	// Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"]
	// Host = "localhost:8000"
	// RemoteAddr = "127.0.0.1:64549"
	// Form["ff"] = ["23"]
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
