package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"sort"
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

	// 先从短到长排序
	sort.SliceStable(config.ArgLint, func(i, j int) bool {
		return len(config.ArgLint[i].X86) > len(config.ArgLint[j].X86)
	})

	// 替换文件内容
	content, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	result := xinchuang_build_lint.ReplaceAndRemove(string(content), config)
	fmt.Println(result)
}
