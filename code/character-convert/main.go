package main

import (
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	f, err := os.Open("my_isotext.txt")
	if err != nil {
		// handle file open error
	}
	out, err := os.Create("my_utf8.txt")
	if err != nil {
		// handler error
	}

	r := charmap.ISO8859_1.NewDecoder().Reader(f)

	io.Copy(out, r)

	out.Close()
	f.Close()
}

//http://technosophos.com/2016/03/09/go-quickly-converting-character-encodings.html
