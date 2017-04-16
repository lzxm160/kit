export GOROOT=/usr/lib/golang
export GOPATH=/root/gokit
export GOBIN=/root/gokit

rm -fr gokit
go build -o gokit myproject/profilesvc
./gokit
