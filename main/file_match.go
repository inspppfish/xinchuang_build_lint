package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"xinchuang_build_lint"
)

func main() {

	// 命令行参数解析
	absPath, cfgFile := xinchuang_build_lint.ArgParse()
	log.Println("absPath:", absPath)
	log.Println("cfgFile:", cfgFile)

	// 解析toml配置文件
	var config xinchuang_build_lint.Config
	if _, err := toml.DecodeFile(cfgFile, &config); err != nil {
		panic(err)
	}

	err := xinchuang_build_lint.Match(absPath, config)
	if err != nil {
		return
	}

}
