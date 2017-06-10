package main

import "expvar"
import "net"
import "fmt"
import "net/http"

var (
    counts = expvar.NewMap("counters")
)

func init() {
    counts.Add("a", 10)
    counts.Add("b", 10)
}

func main() {
    sock, err := net.Listen("tcp", "localhost:8123")
    if err != nil {
        panic("sock error")
    }
    go func() {
        fmt.Println("HTTP now available at port 8123")
        http.Serve(sock, nil)
    }()
    fmt.Println("hello")
    select {}
}
//https://my.oschina.net/moooofly/blog/826460
//https://orangetux.nl/post/expvar_in_action/
