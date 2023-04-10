# 项目名称：xinchuang_build_lint
此文档存在不少gpt生成的部分，会有一些bug
## 简介
xinchuang_build_lint 是一个用于工程化地检索和修改构建配置文件中构建选项的 Go 语言工具。通过此工具，用户可以方便地针对特定文件类型和内容进行自动替换或移除操作。

安装与使用
确保您的计算机上已安装 Go 语言环境。
使用以下命令安装依赖库：
arduino
Copy code
go get -u github.com/BurntSushi/toml
编译 main.go 文件：

```go build main.go```

将生成的二进制文件（例如，main 或 main.exe）添加到您的系统环境变量中，以便在任何位置运行此工具。
运行可执行文件的一个实例:

```main -p=target_build_files -toml=rules/config.toml -r=true```

## 配置
xinchuang_build_lint 通过 TOML 格式的配置文件进行配置。以下是一个配置文件示例：

```
[[filematch]]
pattern = "Makefile"
action = "output"

[[filematch]]
pattern = "CMakeLists.txt"
action = "output"

[[arglint]]
x86 = "-march=i686"
armv8 = "-march=armv8-a"
action = "replace"

[[arglint]]
x86 = "-m32"
armv8 = "-mabi=lp64"
action = "replace"
```
在这个示例中，工具将搜索所有[[filematch]]块指定的的文件，
并将其中的 x86对应的字符串 替换为 armv8对应的字符串。
用户可以根据实际需求添加或修改 [[filematch]] 和 [[arglint]] 配置项。

## 输出
xinchuang_build_lint 将根据配置文件中指定的规则修改所选文件。
工具不会输出详细的替换过程信息，但可以确保修改的准确性。

## 已知限制
xinchuang_build_lint 不具备并发能力。
处理过程中产生的文件名和具体替换过程不够透明。
尽管有以上限制，xinchuang_build_lint 仍是一个有用的工具，
可以帮助您轻松地查找和修改构建配置文件中的构建选项。如果您需要更多的自定义选项或功能，请根据项目需求修改源代码。

## 开发文档(调试项目)
在调试 xinchuang_build_lint 项目时，可以使用 filematch.go，lint.go 和 runall.sh 脚本。这些脚本将帮助您更轻松地测试和调试项目。

### filematch.go
filematch.go 主要用于测试文件匹配功能。它将根据命令行参数和配置文件中的规则查找匹配的文件，并打印出匹配的文件路径。

要运行 filematch.go，请在项目根目录中执行以下命令：

```
go run filematch.go -p=target_build_files -toml=rules/config.toml
```

### lint.go
lint.go 主要用于测试替换和删除文件内容的功能。它将根据命令行参数和配置文件中的规则对给定文件执行替换和删除操作，并输出修改后的文件内容。

要运行 lint.go，请在项目根目录中执行以下命令：

```
go run lint.go -p=target_file -toml=rules/config.toml
```

### runall.sh
runall.sh 是一个 Shell 脚本，用于批量执行 filematch.go 和 lint.go。这个脚本会先运行 filematch.go 来获取匹配的文件列表，然后使用 xargs 并行运行 lint.go 对这些文件执行替换和删除操作。

要运行 runall.sh，请确保您的操作系统支持 Bash（对于 Windows，可以使用 Git Bash 或 WSL），然后在项目根目录中执行以下命令：

```
sh runall.sh
```
注意：在运行 runall.sh 之前，请确保您已经正确设置了命令行参数（如 -p 和 -toml）以及 rules/config.toml 中的配置项。

通过这些脚本，您可以更轻松地调试和测试 xinchuang_build_lint 项目的功能，并根据需要进行修改。
