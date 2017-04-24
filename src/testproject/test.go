package main
import (
	"fmt"
	"time"
	"sync"
	"os"
	"errors"
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
	woffset int64
	roffset int64
	rmutex sync.Mutex
	wmutex sync.Mutex
	dataLen uint32
}
func NewCocurrencyFile(path string,blocksize uint32)(cocurrencyFile,error) {
	f,err:=os.Create(path)
	if err!=nil{
		return nil,err
	}
	if filesize==0{
		return nil,errors.New("invalid size of file")
	}
	df:=&myFile{f:f,dataLen:blocksize}
	return df,nil
}
func (this *myFile)Read()(rsn int64,d []byte,err error) {
	fmt.Println("read")
	this.rmutex.Lock()
	offset:=df.roffset
	df.roffset+=int64(df.dataLen)
	this.rmutex.Unlock()
	rsn=offset/int64(df.dataLen)
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	d=make([]byte,df.dataLen)
	_,err:=df.f.ReadAt(d,offset)
	return
}
func (this *myFile)Write()(wsn int64,err error) {
	return
}
func (this *myFile)RSN()int64 {
	return 0
}
func (this *myFile)WSN()int64 {
	return 0
}
func (this *myFile)Close()error {
	return nil
}
func test() {
	df,err:=NewCocurrencyFile("test.log",3)
	if err!=nil{
		fmt.Println(err)	
	}
	go func() {
		_,d,_=df.Read()
		fmt.Println("a:",d)
	}
	_,d,_=df.Read()
	fmt.Println("b:",d)
}
func main() {
	test()
	time.Sleep(time.Second)
}