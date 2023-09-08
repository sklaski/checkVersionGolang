build:
	go build -o binary/checkVersionGolang_x86_64 cmd/checkVersionGolang/main.go
	env GOOS=linux GOARCH=arm64 go build -o binary/checkVersionGolang_arm64 cmd/checkVersionGolang/main.go
