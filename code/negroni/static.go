package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	// http.FileServer的使用例子, 若你预期要"像伺服器"而非"中间件"的行为
	// mux.Handle("/public", http.FileServer(http.Dir("/home/public")))

	n := negroni.New()
	n.Use(negroni.NewStatic(http.Dir("/tmp")))
	n.UseHandler(mux)

	http.ListenAndServe(":3002", n)
}
