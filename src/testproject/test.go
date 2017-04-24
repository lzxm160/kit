package main
import (
	"fmt"
	"time"
	"sync"
	"os"
)
type cocurrencyFile interface{
	Read()(rsn int64,d []byte,err error)
	Write()(wsn int64,err error)
	RSN()int64
	WSN()int64
	Close()error
}
type myFile struct{
	f *os.File
	fmutex sync.RWMutex
	wsn int64
	rsn int64
	rmutex sync.Mutex
	wmutex sync.Mutex
	dataLen uint32
}
func NewCocurrencyFile(path string,filesize uint32)(*cocurrencyFile,error) {
	f,err:=os.Create(path)
	if err!=nil{
		return nil,err
	}
	if filesize==0{
		return nil,error.New("invalid size of file")
	}
	df:=&myFile{f:f,dataLen:filesize}
	return df,nil
}
func (this *myFile)Read()(rsn int64,d []byte,err error) {
	fmt.Println("read")
}
func (this *myFile)Write()(wsn int64,err error) {
	
}
func (this *myFile)RSN()int64 {
	
}
func (this *myFile)WSN()int64 {
	
}
func (this *myFile)Close()error {
	
}
func test() {
	df,err:=NewCocurrencyFile("test.log",10000)
	if err!=nil{
		df.Read()
	}
}
func main() {
	test()
	time.Sleep(time.Second)
}