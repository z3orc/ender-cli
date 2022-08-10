BINARY_NAME=ender-cli

run:
	./ender-cli.exe

build:
	go build -o build/ender-cli.exe

compile:
	GOARCH=amd64 GOOS=darwin go build -o build/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o build/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=window go build -o build/${BINARY_NAME}-windows main.go

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows