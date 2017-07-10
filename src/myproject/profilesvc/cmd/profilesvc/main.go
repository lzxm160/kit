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
	"testing"
)
var bytePool=sync.Pool{
	New:newPool,
}
func newPool() interface{}{
	b:=make([]byte,1024)
	return &b
}
func BenchmarkAlloc(b *testing.B) {
	for i:=0;i<b.N;i++{
		o:=make([]byte,1024)
		_=o
	}
}
func BenchmarkPool(b *testing.B) {
	for i:=0;i<b.N;i++{
		o:=bytePool.Get().(*[]byte)
		_=o
		bytePool.Put(o)
	}
}
func main() {
	var mutex sync.RWMutex
	for i:=0;i<3;i++{
		go func(i int) {
			fmt.Println("try read lock ",i)
			mutex.RLock()
			fmt.Println("locked ",i)
			time.Sleep(time.Second)
			fmt.Println("try read unlock ",i)
			mutex.RUnlock()
			fmt.Println("unlocked ",i)
		}(i)
	}
	time.Sleep(time.Millisecond*100)
	fmt.Println("try to write lock main")
	mutex.Lock()
	fmt.Println("write locked")
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
// }
