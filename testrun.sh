export GOROOT=/usr/lib/golang
export GOPATH=/root/gopath:`pwd`
export GOBIN=

#rm -fr gokit
#go build -o gokit myproject/profilesvc/cmd/profilesvc
#./gokit
#go run src/testproject/test.go
#go test -bench myproject/profilesvc/cmd/profilesvc/main.go
#go build -o gokit myproject/profilesvc/cmd/profilesvc
cd src/myproject/profilesvc/cmd/profilesvc/
#go test -bench .
go test -race .