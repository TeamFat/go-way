package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/toolkits/file"
)

var (
	cfg = flag.String("cfg", "config.json", "cfg")
)

type ServerConfifg struct {
	Url string
}

func init() {
	flag.Parse()
}

func ParseConfig(cfg string) ServerConfifg {

	configContent, _ := file.ToTrimString(cfg)
	var c ServerConfifg
	json.Unmarshal([]byte(configContent), &c)
	return c

}

func main() {
	c := ParseConfig(*cfg)
	fmt.Println(c.Url)
	os.Exit(0)
}
