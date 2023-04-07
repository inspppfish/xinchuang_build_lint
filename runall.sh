

go run main/file_match.go -p=target_build_files -toml=rules/config.toml | xargs -I {} sh -c "go run main/lint.go -p={} -toml=rules/config.toml > {}.new"