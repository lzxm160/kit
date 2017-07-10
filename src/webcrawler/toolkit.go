package main
import (

)

type Buffer interface{
	Cap()uint32
	Len()uint32
	Put(data interface{})(bool,error)
	Get()(interface{},error)
	Close()bool
	Closed()bool
}
type myBuffer struct{
	ch chan interface{}
	closed uint32
	closingLock sync.RWMutex
}
func (this *myBuffer)Cap()uint32{
	return cap(this.ch)
}
func (this *myBuffer)Len()uint32{
	return len(this.ch)
}
func (this *myBuffer)Put(data interface{})(ok bool,err error){
	this.closingLock.RLock()
	defer this.closingLock.RUlock()
	if this.Closed(){
		return false,ErrClosedBuffer
	}
	select{
	case this.ch<-data:
		ok=true
	case default:
		ok=false
	}
	return
}
func (this *myBuffer)Get()(interface{},error){
	select{
	case data,ok:=<-this.ch:
		if !ok{
			return nil,ErrClosedBuffer
		}
		return data,nil
	default:
		return nil,nil
	}
}
var ErrClosedBuffer=errors.New("error closed buffer")
var ErrClosedBufferPool=errors.New("error closed buffer pool")
func (this *myBuffer)Close()bool{
	if atomic.CompareAndSwapUint32(&this.closed,0,1){
		this.closingLock.Lock()
		close(this.ch)
		this.closingLock.Unlock()
		return true
	}
	return false
}
func (this *myBuffer)Closed()bool{
	if atomic.LoadUint32(&this.closed)==0{
		return false
	}

	return true
}
func NewBuffer(size uint32)(Buffer,error) {
	if size==0{
		return nil,errors.New(fmt.Sprintf("illegal size:%d",size))
	}
	return &myBuffer{ch:make(chan interface{},size),closed:0},nil
}

////////////////////////////////////////////////
type Pool interface{
	BufferCap() uint32
	MaxBufferNumber() uint32
	BufferNumber() uint32
	Total() uint64
	Put(data interface{})error
	Get()(interface{},error)
	Close()bool
	Closed()bool
}

type myPool struct{
	bufferCap uint32
	maxBufferNumber uint32
	bufferNumber uint32
	total uint64
	bufCh chan Buffer
	closed uint32
	rwlock sync.RWMutex
}
func NewPool()Pool {
	
}
func (this *myPool)BufferCap() uint32{

}
func (this *myPool)MaxBufferNumber() uint32{

}
func (this *myPool)BufferNumber() uint32{

}
func (this *myPool)Total() uint64{

}
func (this *myPool)Put(datum interface{})(err error){
	if this.Closed(){
		return ErrClosedBufferPool
	}
	var count uint32
	maxCount:=this.BufferNumber()*5
	var ok bool
	for buf:=range this.bufCh{
		ok,err=this.putData(buf,datum,&count,maxCount)
		if ok||err!=nil{
			break
		}
	}
	return
}
func (this *myPool)putData(buf Buffer,datum interface{},count *uint32,maxCount uint32)(ok bool,err error) {
	if this.Closed(){
		return false,ErrClosedBufferPool
	}
	defer func() {
		this.rwlock.RLock()
		if this.Closed(){
			atomic.AddUint32(&this.bufferNumber,^uint32(0))
			err=ErrClosedBufferPool
		}else{
			this.bufCh<-buf
		}
		this.rwlock.RUlock()
	}()
	ok,err=buf.Put(datum)
	if ok{
		atomic.AddUint64(&this.total,1)
		return
	}
	if err!=nil{
		return
	}
	(*count)++
	if *count>=maxCount&&this.BufferNumber()<this.MaxBufferNumber(){
		this.rwlock.Lock()
		if this.BufferNumber()<this.MaxBufferNumber(){
			if this.Closed(){
				this.rwlock.Unlock()
				return 
			}
			newBuf,_:=NewBuffer(this.bufferCap)
			newBuf.Put(datum)
			this.bufCh<-newBuf
			atomic.AddUint32(&this.bufferNumber,1)
			atomic.AddUint64(&this.total,1)
			ok=true
		}
		this.rwlock.Unlock()
		*count=0//貌似没有必要
	}
	return
}
func (this *myPool)Get()(interface{},error){
	if this.Closed(){
		return ErrClosedBufferPool
	}
	var count uint32
	maxCount:=this.BufferNumber()*5
	var datum interface{}
	for buf:=range this.bufCh{
		datum,err=this.getData(buf,&count,maxCount)
		if datum!=nil||err!=nil{
			break
		}
	}
	return
}
func (this *myPool)getData(buf Buffer,count *uint32,maxCount uint32)(datum interface{},err error) {
	if this.Closed(){
		return nil,ErrClosedBufferPool
	}
	defer func() {
		if *count>maxCount&&buf.Len()==0&&this.BufferNumber()>1{
			buf.Close()
			atomic.AddUint32(&this.bufferNumber,^uint32(0))
			*count=0
			return
		}
		this.rwlock.RLock()
		if this.Closed(){
			atomic.AddUint32(&this.bufferNumber,^uint32(0))
			err=ErrClosedBufferPool
		}else{
			this.bufCh<-buf
		}
		this.rwlock.RUlock()
	}()
	datum,err=buf.Gut()
	if datum!=nil{
		atomic.AddUint64(&this.total,^uint64(0))
		return
	}
	if err!=nil{
		return
	}
	(*count)++
	
	return
}
func (this *myPool)Close()bool{
	if !atomic.CompareAndSwapUint32(&this.closed,0,1){
		return false
	}
	this.rwlock.Lock()
	defer this.rwlock.Unlock()
	close(this.bufCh)
	for buf:=range this.bufCh{
		buf.Close()
	}
	return true
}
func (this *myPool)Closed()bool{

}
////////////////////////////////////////////
type MultipleReader interface{
	Reader()io.ReadCloser
}
type myMultipleReader struct{
	data []byte
}
func NewMultipleReader(reader io.Reader)(MultipleReader,error) {
	var data[]byte
	var err error
	if reader!=nil{
		data,err=ioutil.ReadAll(reader)
		if err!=nil{
			return nil,fmt.Errorf("multiple reader:cant create new one:%s",err)
		}else{
			data=[]byte{}
		}
		return &myMultipleReader{data:data},nil

	}
}
func (this *myMultipleReader)Reader()io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(this.data))
}