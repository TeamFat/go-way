package main

import (
	"fmt"
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
	recovery.PanicHandlerFunc = reportToSentry
	n.Use(recovery)
	n.UseHandler(mux)

	http.ListenAndServe(":3003", n)
}

func reportToSentry(info *negroni.PanicInformation) {
	// 在这写些程式回报错误给Sentry
	fmt.Println(info)
}
