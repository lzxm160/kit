package main
import (
	"fmt"
	"time"
	"sync"
	"os"
	"errors"
	// "reflect"
)
type cocurrencyFile interface{
	Read()(rsn int64,d []byte,err error)
	Write(d []byte)(wsn int64,err error)
	RSN()int64
	WSN()int64
	Close()error
	Roffset()int64 
	Woffset()int64 
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
	if blocksize==0{
		return nil,errors.New("invalid size of file")
	}
	df:=&myFile{f:f,dataLen:blocksize}
	return df,nil
}
func (this *myFile)Read()(rsn int64,d []byte,err error) {
	fmt.Println("read")
	this.rmutex.Lock()
	offset:=this.roffset
	this.roffset+=int64(this.dataLen)
	this.rmutex.Unlock()
	rsn=offset/int64(this.dataLen)
	this.fmutex.Lock()
	defer this.fmutex.Unlock()
	d=make([]byte,this.dataLen)
	_,err=this.f.ReadAt(d,offset)
	return
}
func (this *myFile)Write(d []byte)(wsn int64,err error) {
	fmt.Println("write")
	this.wmutex.Lock()
	offset:=this.woffset
	woffsetadd:=0
	if len(d)>this.dataLen{
		woffsetadd=this.dataLen
	}else{
		woffsetadd=len(d)
	}
	this.woffset+=int64(woffsetadd)
	this.wmutex.Unlock()
	wsn=offset/int64(this.dataLen)

	var bytes []byte
	if len(d)>int(this.dataLen){
		bytes=d[0:this.dataLen]
	}else{
		bytes=d
	}
	this.fmutex.Lock()
	defer this.fmutex.Unlock()
	
	fmt.Println("len d:",len(bytes))
	_,err=this.f.Write(bytes)
	return

}
func (this *myFile)RSN()int64 {
	this.rmutex.Lock()
	defer this.rmutex.Unlock()

	return this.roffset/int64(this.dataLen)
}
func (this *myFile)WSN()int64 {
	this.wmutex.Lock()
	defer this.wmutex.Unlock()
	return this.woffset/int64(this.dataLen)
}
func (this *myFile)Close()error {
	return this.f.Close()
}
func (this *myFile)Roffset()int64 {
	return this.roffset
}
func (this *myFile)Woffset()int64 {
	return this.woffset
}
func test() {
	df,err:=NewCocurrencyFile("test.log",3)
	if err!=nil{
		fmt.Println(err)	
	}
	//////
	// df.(type).f.write([]byte{1,2,3,4,5,6,7,8,9,0})
	// v := reflect.ValueOf(&df)
	// v.Interface().(myFile).f.Write([]byte{1,2,3,4,5,6,7,8,9,0})
	syncchan:=make(chan struct{},2)
	go func() {
		df.Write([]byte{1,2,3})
		syncchan<-struct{}{}
	}()
	go func() {
		df.Write([]byte{4,5})
		syncchan<-struct{}{}
	}()
	<-syncchan
	<-syncchan
	go func() {
		_,d,_:=df.Read()
		fmt.Println("a:",d)
		syncchan<-struct{}{}
	}()
	_,d,_:=df.Read()
	<-syncchan
	fmt.Println("b:",d)

	// fmt.Println(df.Roffset())
	fmt.Println(df.RSN())

	// fmt.Println(df.Woffset())
	fmt.Println(df.WSN())	

	df.Close()
	// v := reflect.ValueOf(&df)	
	// fmt.Println(v.Interface().(myFile).Woffset())
	// fmt.Println(v.Interface().(myFile).Roffset())
	// v0 := make([]reflect.Value, 0)
	// v.MethodByName("Woffset").Call(v0)
	fmt.Println(df.Roffset())
	fmt.Println(df.Woffset())
}
func main() {
	test()
	time.Sleep(time.Second)
}