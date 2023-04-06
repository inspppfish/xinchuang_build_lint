package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

func main() {
	var config Config
	if _, err := toml.DecodeFile("rules/build_file.toml", &config); err != nil {
		panic(err)
	}
	fmt.Println(config)
}
