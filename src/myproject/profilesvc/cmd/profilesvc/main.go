package main

import (
	_"flag"
	"fmt"
	_"net/http"
	_"os"
	_"os/signal"
	_"syscall"
	"sync"
	"time"
	_"github.com/go-kit/kit/examples/profilesvc"
	_"github.com/go-kit/kit/log"
	_"testing"
	"sync/atomic"
)
func test1() {
	//var mutex=new(sync.Mutex)
	var mutex sync.Mutex
	cond:=sync.NewCond(&mutex)
	done:=false
	cond.L.Lock()
	go func() {
		time.Sleep(time.Second*3)
		done=true
		cond.Signal()
	}()
	if(!done){
		cond.Wait()
	}
	fmt.Println("done")
}
func addValueAtomic(delta int32) {
	var value int32=3

	for{
		v:=atomic.LoadInt32(&value)
		fmt.Println("1:",v)
		if atomic.CompareAndSwapInt32(&value,v,v+delta){
			break
		}
	}
	fmt.Println("2:",value)
}
func test2() {
	// var u64 uint64
	// u64=3
	// atomic.AddUint64(&u64,^uint64(3-1))
	// fmt.Println(u64)
	addValueAtomic(2)
}
func test3() {
	var atomicVal atomic.Value
	atomicVal.Store([]int{1,3,5,7})
	another(atomicVal)
	fmt.Printf("%+v",atomicVal)
}
func another(c atomic.Value) {
	c.Store([]int{2,4,6,8})
}
type ConcurrentArrayInterface interface{
	Set(index uint32,elem int)(err error)
	Get(index uint32)(elem int,err error)
	Len()uint32
}
type ConcurrentArray struct{
	length uint32
	val atomic.Value
}
func NewConcurrencyArray(len uint32)ConcurrentArrayInterface {
	arr:=ConcurrentArray{}
	arr.length=len
	arr.val.Store(make([]int,arr.length))
	return &arr
}
func (this *ConcurrentArray)Set(index uint32,elem int)(err error) {
	newArray:=make([]int,this.length)
	copy(newArray,this.val.Load().([]int))
	newArray[index]=elem
	this.val.Store(newArray)
	return
}
func (this *ConcurrentArray)Get(index uint32) (val int,err error){
	val=this.val.Load().([]int)[index]
	return
}
func (this *ConcurrentArray)Len() (len uint32){
	len=this.length
	return
}
func test() {
	te:=NewConcurrencyArray(5)
	for i:=0;i<5;i++{
		go func(i int) {
			te.Set(uint32(i),uint32(i)*uint32(i))
		}()	
	}
	for i:=0;i<5;i++{
		v,e:=te.Get(i)
		if(!e){
			fmt.Println(v)
		}
	}
}
func main() {
	test()
	// var once sync.Once
	// onceFunc:=func() {
	// 	fmt.Println("once")
	// }
	// for i:=0;i<10;i++{
	// 	go func() {
	// 		once.Do(onceFunc)
	// 	}()
	// }
	// time.Sleep(time.Second*3)
// 	var mutex sync.RWMutex
// 	for i:=0;i<3;i++{
// 		go func(i int) {
// 			fmt.Println("try read lock ",i)
// 			mutex.RLock()
// 			fmt.Println("locked ",i)
// 			time.Sleep(time.Second)
// 			fmt.Println("try read unlock ",i)
// 			mutex.RUnlock()
// 			fmt.Println("unlocked ",i)
// 		}(i)
// 	}
// 	time.Sleep(time.Millisecond*100)
// 	fmt.Println("try to write lock main")
// 	mutex.Lock()
// 	fmt.Println("write locked")
	// var (
	// 	httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	// )
	// flag.Parse()

	// var logger log.Logger
	// {
	// 	logger = log.NewLogfmtLogger(os.Stderr)
	// 	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	// 	logger = log.With(logger, "caller", log.DefaultCaller)
	// }

	// var s profilesvc.Service
	// {
	// 	s = profilesvc.NewInmemService()
	// 	s = profilesvc.LoggingMiddleware(logger)(s)
	// }

	// var h http.Handler
	// {
	// 	h = profilesvc.MakeHTTPHandler(s, log.With(logger, "component", "HTTP"))
	// }

	// errs := make(chan error)
	// go func() {
	// 	c := make(chan os.Signal)
	// 	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	// 	errs <- fmt.Errorf("%s", <-c)
	// }()

	// go func() {
	// 	logger.Log("transport", "HTTP", "addr", *httpAddr)
	// 	errs <- http.ListenAndServe(*httpAddr, h)
	// }()

	// logger.Log("exit", <-errs)
}
