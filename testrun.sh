export GOROOT=/usr/lib/golang
export GOPATH=`pwd`;`pwd`/src
export GOBIN=`pwd`

#rm -fr gokit
#go build -o gokit myproject/profilesvc/cmd/profilesvc
#./gokit
#go run src/testproject/test.go
#go test -bench myproject/profilesvc/cmd/profilesvc/main.go
go test -v myproject/profilesvc/cmd/profilesvc/main.go