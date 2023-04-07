package xinchuang_build_lint

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type MakeFileName string

/* 字符串数组存储Makefile的名字，可以自行添加 */
var MakeFileNames = []MakeFileName{"Makefile", "makefile", "CMakeLists.txt", "BUILD", ".bazelrc"}

// 定义替换规则
var rules = map[string]string{
	"-march=i686": "-march=armv8.2-a",
	"-march":      "-march=armv8.2-a",
	"-m32":        "-mabi=lp64",
	// 这里可以添加更多的x替换规则
}

/* 以程序当前所在路径为源路径开始遍历 */
func Init() string {
	root, err := os.Getwd()
	fmt.Println()
	fmt.Println("now root is " + root)
	fmt.Println("#############################################")
	if err != nil {
		fmt.Println(err)
		return root
	}

	return root
}

/* 匹配文件 */
func MatchFile(fileChan chan string, isdoneChan chan bool) {
	for path := range fileChan {
		for _, fileName := range MakeFileNames { // 检测文件名
			if ok, _ := filepath.Match(string(fileName), filepath.Base(path)); ok {
				fmt.Println("path: " + path)
				ReadAndSubFileLine(path)
				break
			}
		}
	}
	isdoneChan <- true
}

/* 遍历文件夹 */
func TraverseDir(root string, fileChan chan string) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fileChan <- path
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

/* 读取文件里的每一行 */
func ReadAndSubFileLine(path string) {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	fmt.Println("############### FileLine Start ###############")
	scanner := bufio.NewScanner(file)
	var offset int64
	for scanner.Scan() {
		// 读取一行
		line := scanner.Text()

		// 依次应用每个替换规则
		for oldStr, newStr := range rules {
			line = strings.ReplaceAll(line, oldStr, newStr)
		}

		// 计算偏移量并将替换后的行写回文件
		lineBytes := []byte(line + "\n")
		if _, err := file.WriteAt([]byte(line+"\n"), offset); err != nil {
			panic(err)
		}
		offset += int64(len(lineBytes))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println("############### FileLine Finish ###############")
	fmt.Println()
}

//func main() {
//
//	root := Init()
//
//	/* 用两个channel + goroutine 的方式并发，保证在较大数据量情况下可以运行 */
//	fileChan := make(chan string)
//	isdoneChan := make(chan bool)
//
//	/* 用 goruntine 进行文件匹配 */
//	go func() {
//		MatchFile(fileChan, isdoneChan)
//	}()
//
//	/* 并发进行 root 下的遍历，使用匿名函数把 fileChan 内嵌进去 */
//	TraverseDir(root, fileChan)
//
//	close(fileChan)
//}
