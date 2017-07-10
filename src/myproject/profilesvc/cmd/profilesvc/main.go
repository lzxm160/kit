package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"sync"
	"time"
	"github.com/go-kit/kit/examples/profilesvc"
	"github.com/go-kit/kit/log"
)

func main() {
	var mutex sync.Mutex
	fmt.Println("start lock main")
	mutex.Lock()
	fmt.Println("main is locked")
	for i:=1;i<=3;i++{
		go func(int i) {
			fmt.Println("go start lock %d",i)
			mutex.Lock()
			fmt.Println("go lock is locked %d",i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("start unlock main")
	mutex.Unlock()
	fmt.Println("unlocked main")
	
	time.Sleep(time.Second*3)
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
