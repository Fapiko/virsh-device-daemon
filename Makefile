build:
	GOOS=linux GOARCH=amd64 go build github.com/fapiko/virsh-device-daemon
	GOOS=windows GOARCH=amd64 go build -o virsh-device-daemon-win64.exe github.com/fapiko/virsh-device-daemon
