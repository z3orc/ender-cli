BINARY_NAME=ender-cli

run:
	./ender-cli.exe

compile:
	GOARCH=amd64 GOOS=darwin go build -o build/${BINARY_NAME}-darwin-x64 main.go
	GOARCH=arm64 GOOS=darwin go build -o build/${BINARY_NAME}-darwin-arm64 main.go
	GOARCH=amd64 GOOS=linux go build -o build/${BINARY_NAME}-linux-x64 main.go
	GOARCH=arm64 GOOS=linux go build -o build/${BINARY_NAME}-linux-arm64 main.go
	# GOARCH=amd64 GOOS=windows go build -o build/${BINARY_NAME}-windows-x64.exe main.go

compile-windows:
	go build -o build/${BINARY_NAME}-windows-x64.exe

clean:
	go clean
	rm ${BINARY_NAME}-darwin-x64
	rm ${BINARY_NAME}-darwin-arm64
	rm ${BINARY_NAME}-linux-x64
	rm ${BINARY_NAME}-linux-arm64
	# rm ${BINARY_NAME}-windows-x64.exe