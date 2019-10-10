if [ -e ./server.pid ]; then
    kill -2 `cat ./server.pid`
    rm -f ./server.pid ./server.sock
fi
GOOS=linux GOARCH=amd64 go build -o server src/main.go
