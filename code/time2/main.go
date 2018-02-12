package main

import (
	"fmt"
	"strings"
)

func Substr(str string, start int, end int) string {
	rs := []byte(str)
	length := len(rs)

	if start < 0 || start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}
	return string(rs[start:end])
}
func main() {
	time := "2018-02-11T14:42:40"
	new := strings.Replace(time, "T", " ", -1)

	fmt.Println(Substr(new, 0, 16))
}
