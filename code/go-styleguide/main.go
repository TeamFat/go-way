package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	//https://github.com/bahlo/go-styleguide
	_, err := os.Open("foo.txt")
	if err != nil {
		fmt.Println(errors.Wrap(err, "open foo.txt failed"))
	}
}
