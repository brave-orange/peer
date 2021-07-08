linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o udpHoleServer server.go
	chmod +x udpHoleServer
osx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o udpHoleServer server.go
	chmod +x udpHoleServer