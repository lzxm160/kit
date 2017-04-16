export GOPATH=/root/gokit
export GOBIN=/root/gokit

rm -fr gokit
go build -o gokit myproject/stringsvc3
./gokit
