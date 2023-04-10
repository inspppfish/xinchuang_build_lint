package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"xinchuang_build_lint"
)

func main() {

	// 命令行参数解析
	absPath, recursion, cfgFile := xinchuang_build_lint.ArgParse()
	log.Println("absPath:", absPath)
	log.Println("cfgFile:", cfgFile)
	log.Println("recursion:", recursion)

	// 解析toml配置文件
	var config xinchuang_build_lint.Config
	if _, err := toml.DecodeFile(cfgFile, &config); err != nil {
		panic(err)
	}

	result, err := xinchuang_build_lint.Match(absPath, recursion, config)
	if err != nil {
		panic(err)
	}
	for _, path := range result {
		fmt.Println(path)
	}
}
