export GOROOT=/usr/lib/golang
export GOPATH=`pwd`
export GOBIN=`pwd`

rm -fr gokit
go build -o gokit myproject/profilesvc/cmd/profilesvc
./gokit
