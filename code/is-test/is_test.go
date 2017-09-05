package main

import (
	"testing"

	"github.com/matryer/is"
)

func Test(t *testing.T) {
	is := is.New(t)
	a, b := 1, 1
	is.Equal(a, b) // expect to be the same
}
