package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"sort"
	"xinchuang_build_lint"
)

func main() {
	// 命令行参数解析
	absPath, recursion, cfgFile := xinchuang_build_lint.ArgParse()
	log.Println("workPath:", absPath)
	log.Println("cfgFile:", cfgFile)

	// 解析toml配置文件
	var config xinchuang_build_lint.Config
	if _, err := toml.DecodeFile(cfgFile, &config); err != nil {
		panic(err)
	}

	// 匹配文件名
	filenames, err := xinchuang_build_lint.Match(absPath, recursion, config)
	if err != nil {
		panic(err)
	}

	// 先把被匹配的lint目标从长到短排序
	sort.SliceStable(config.ArgLint, func(i, j int) bool {
		return len(config.ArgLint[i].X86) > len(config.ArgLint[j].X86)
	})

	// 替换文件内容
	for _, filename := range filenames {
		file, _ := os.OpenFile(filename, os.O_RDWR, 0666)
		content, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		result := xinchuang_build_lint.ReplaceAndRemove(string(content), config)
		_, err = file.Write([]byte(result))
		if err != nil {
			panic(err)
		}
	}
}
