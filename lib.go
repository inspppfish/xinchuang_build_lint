package xinchuang_build_lint

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	FileMatch []struct {
		Pattern string
		Action  string
	}
	ArgLint []struct {
		X86    string
		Armv8  string
		Action string
	}
}

// ArgParse 命令行参数解析
func ArgParse() (string, string) {
	// 获取命令行参数 -p 并设定默认值
	defaultDir, _ := os.Getwd()
	pathFlag := flag.String("p", defaultDir, "help massage for -p : the path that the match program works")
	tomlFlag := flag.String("toml", "rules/config.toml", "help massage for -toml : the toml config file that the match program works")

	// 解析命令行参数，应该在所有flag声明之后
	flag.Parse()

	// 绝对路径，输出到log避免被重定向
	absPath, _ := filepath.Abs(*pathFlag)
	cfgFile, _ := filepath.Abs(*tomlFlag)
	return absPath, cfgFile

}

// Match 根据设置匹配文件并输出
// todo: 错误处理看起来很丑
func Match(path string, config Config) error {
	for _, match := range config.FileMatch {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// 匹配文件名, 根据action决定是否输出文件路径
			if ok, _ := filepath.Match(match.Pattern, filepath.Base(path)); ok && match.Action == "output" {
				fmt.Println(path)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// ReplaceAndRemove 替换和删除文件内容并输出
// todo:修改替换算法，避免重复替换
func ReplaceAndRemove(content string, config Config) error {
	// 先不用手工的方法，再找找有没有go的替换算法
	// mask := make([]bool, len(content))
	for _, lint := range config.ArgLint {
		// 替换文件内容
		content = strings.ReplaceAll(content, lint.X86, lint.Armv8)
		// 删除文件内容
		content = strings.ReplaceAll(content, lint.X86, "")
	}
	fmt.Println(content)
	return nil
}
