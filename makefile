BINARY_NAME=http-recorder

default: build

build:
	go build main.go

builddarwin:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin-amd64 main.go
	tar -zcvf ${BINARY_NAME}-darwin-amd64.tar.gz ${BINARY_NAME}-darwin-amd64

buildlinux:
	GOARCH=amd64 GOOS=linux  go build -o ${BINARY_NAME}-linux-amd64 main.go
	tar -zcvf ${BINARY_NAME}-linux-amd64.tar.gz ${BINARY_NAME}-linux-amd64

buildwindows:
	GOARCH=amd64 GOOS=window go build -o ${BINARY_NAME}-windows-amd64.exe main.go
	tar -zcvf ${BINARY_NAME}-windows-amd64.tar.gz ${BINARY_NAME}-windows-amd64

buildall:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin-amd64 main.go
	GOARCH=amd64 GOOS=linux  go build -o ${BINARY_NAME}-linux-amd64 main.go
	GOARCH=amd64 GOOS=window go build -o ${BINARY_NAME}-windows-amd64.exe main.go
