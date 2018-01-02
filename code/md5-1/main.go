package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "welcome to jb51.net")
	fmt.Printf("%x", h.Sum(nil))
}
