package main

import (
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		panic("oh no")
	})

	n := negroni.New()
	recovery := negroni.NewRecovery()
	recovery.Formatter = &negroni.HTMLPanicFormatter{}
	n.Use(recovery)
	n.UseHandler(mux)

	http.ListenAndServe(":3003", n)
}
