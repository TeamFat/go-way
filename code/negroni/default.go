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

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
	//https://github.com/urfave/negroni/blob/master/translations/README_zh_CN.md
}
