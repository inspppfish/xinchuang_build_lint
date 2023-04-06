package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path/filepath"
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

func main() {

	// 获取命令行参数 -p 并设定默认值
	defaultDir, _ := os.Getwd()
	pathFlag := flag.String("p", defaultDir, "help massage for -p : the path that the match program works")
	tomlFlag := flag.String("toml", "rules/build_file.toml", "help massage for -toml : the toml config file that the match program works")

	// 解析命令行参数，应该在所有flag声明之后
	flag.Parse()

	// 绝对路径，输出到log避免被重定向
	absPath, _ := filepath.Abs(*pathFlag)
	cfgFile, _ := filepath.Abs(*tomlFlag)
	log.Println("absPath:", absPath)
	log.Println("cfgFile:", cfgFile)

	var config Config
	if _, err := toml.DecodeFile(cfgFile, &config); err != nil {
		panic(err)
	}
	//fmt.Println(config)

	//entries, _ := os.ReadDir(absPath)
	//for _, entry := range entries {
	//	// 获取文件或子目录的信息
	//	info, err := entry.Info()
	//	if err != nil {
	//		fmt.Println("Error:", err)
	//		continue
	//	}
	//
	//	// 判断是否为目录
	//	if info.IsDir() {
	//		fmt.Println("Directory:", entry.Name())
	//	} else {
	//		fmt.Println("File:", entry.Name())
	//	}
	//}
}
