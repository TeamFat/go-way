package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/toolkits/file"
)

//go:generate ldetool generate --package main Line.lde
func main() {
	//go run main.go Line_lde.go
	l := &Line{}
	configContent, _ := file.ToTrimString("log")
	scanner := bufio.NewScanner(strings.NewReader(configContent))
	for scanner.Scan() {
		ok, err := l.Extract(scanner.Bytes())
		if !ok {
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		fmt.Println(string(l.Time))
	}
}
