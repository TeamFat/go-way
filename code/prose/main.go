package main

import (
	"fmt"
	"strings"

	"github.com/jdkato/prose/transform"
)

func main() {
	text := "the last of the mohicans"
	tc := transform.NewTitleConverter(transform.APStyle)
	fmt.Println(strings.Title(text)) // The Last Of The Mohicans
	fmt.Println(tc.Title(text))      // The Last of the Mohicans
}
