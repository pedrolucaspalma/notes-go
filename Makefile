build-windows: 
	GOOS=windows GOARCH=amd64 go build -o notes.exe cmd/tui/main.go
build-linux: 
	GOOS=linux GOARCH=amd64 go build -o notes cmd/tui/main.go
build-mac:
	GOOS=darwin GOARCH=arm64 go build -o notes cmd/tui/main.go
