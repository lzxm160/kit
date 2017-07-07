export GOROOT=/usr/lib/golang
export GOPATH=./;./src
export GOBIN=./

rm -fr gokit
go build -o gokit myproject/profilesvc/cmd/profilesvc
./gokit
