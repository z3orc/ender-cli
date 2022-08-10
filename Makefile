BINARY_NAME=ender-cli

run:
	./ender-cli.exe

build:
	go build -o build/ender-cli.exe

compile:
	GOARCH=amd64 GOOS=darwin go build -o build/${BINARY_NAME}-darwin-x64 main.go
	GOARCH=arm64 GOOS=darwin go build -o build/${BINARY_NAME}-darwin-arm64 main.go
	GOARCH=amd64 GOOS=linux go build -o build/${BINARY_NAME}-linux-x64 main.go
	GOARCH=arm64 GOOS=linux go build -o build/${BINARY_NAME}-linux-arm64 main.go
	GOARCH=amd64 GOOS=window go build -o build/${BINARY_NAME}-windows-x64 main.go

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows