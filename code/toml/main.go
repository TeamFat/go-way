package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   string
	Port     int
	Email    string
	Password string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}

var config = Config{}

func init() {
	config.Read()
}

func main() {
	fmt.Println(config)
}
