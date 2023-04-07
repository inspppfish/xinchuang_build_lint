
# todo: 两条平行路径： 1.混淆这个脚本,保留原先架构 2.编写一个新的go函数入口，使得之前的go文件可以编译为一个可执行文件
go run main/file_match.go -p=target_build_files -toml=rules/config.toml | xargs -I {} -P 4 sh -c "go run main/lint.go -p={} -toml=rules/config.toml"