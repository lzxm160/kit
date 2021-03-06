package main
import (
	"fmt"
	"time"
	"sync"
)
func test1() {
	content:=make(chan string,3)
	sync1:=make(chan struct{},1)
	sync2:=make(chan struct{},2)
	go func() {
		sync1<-struct{}{}
		for{
			if elem,ok:=<-content;ok{
				fmt.Println("sync1:",elem)
				// time.Sleep(time.Second)
			}else{
				fmt.Println("sync1 break")
				break
			}
		}
		defer func(){sync2<-struct{}{}}()
	}()
	go func() {
		fmt.Println("sync2")
		str:=[]string{"a","b","c","d"}
		for _,elem:=range str{
			content<-elem
			if elem=="c"{
				<-sync1
			}
		}
		close(content)
		defer func(){sync2<-struct{}{}}()
	}()
	
	<-sync2

	<-sync2
} 
func test2() {
	countmap:=make(chan map[string]int,1)
	sync2:=make(chan struct{},2)
	go func() {
		for{
			if elem,ok:=<-countmap;ok{
				elem["count"]++
				fmt.Println("sync1:",elem)
			}else{
				fmt.Println("sync1 break")
				break
			}
		}
		defer func(){sync2<-struct{}{}}()
	}()
	go func() {
		content:=make(map[string]int)
		content["count"]=0
		for i:=0;i<5;i++{
			countmap<-content
			fmt.Println("sync2:",content)
		}
		close(countmap)
		defer func(){sync2<-struct{}{}}()
	}()
	
	<-sync2

	<-sync2
} 
type Counter struct{
	count int
}
func (this *Counter)String()string {
	return fmt.Sprintf("{count:%d}",this.count)
}

func test3() {
	countmap:=make(chan map[string]*Counter,1)
	sync2:=make(chan struct{},2)
	go func() {
		for{
			if elem,ok:=<-countmap;ok{
				c:=elem["count"]
				c.count++
				fmt.Println("sync1:",elem)
			}else{
				fmt.Println("sync1 break")
				break
			}
		}
		defer func(){sync2<-struct{}{}}()
	}()
	go func() {
		content:=map[string]*Counter{"count":&Counter{}}
		for i:=0;i<5;i++{
			countmap<-content
			fmt.Println("sync2:",content)
		}
		close(countmap)
		defer func(){sync2<-struct{}{}}()
	}()
	
	<-sync2

	<-sync2
} 
func receive(countmap <-chan map[string]*Counter,sync2 chan<- struct{}) {
	// for{
	// 	if elem,ok:=<-countmap;ok{
	// 		c:=elem["count"]
	// 		c.count++
	// 		fmt.Println("sync1:",elem)
	// 	}else{
	// 		fmt.Println("sync1 break")
	// 		break
	// 	}
	// }
	for elem:=range countmap{
		c:=elem["count"]
		c.count++
		fmt.Println("sync1:",elem)
	}
	sync2<-struct{}{}
}
func sends(countmap chan<- map[string]*Counter,sync2 chan<- struct{}) {
	content:=map[string]*Counter{"count":&Counter{}}
	for i:=0;i<5;i++{
		countmap<-content
		fmt.Println("sync2:",content)
	}
	close(countmap)
	sync2<-struct{}{}
}
func test4() {
	countmap:=make(chan map[string]*Counter,1)
	sync2:=make(chan struct{},2)
	go receive(countmap,sync2)
	go sends(countmap,sync2)
	
	<-sync2

	<-sync2
} 
func get_number(i int)int {
	fmt.Println("get_number:",i)
	return num[i]
}
func get_chan(i int)chan int {
	fmt.Println("get_chan:",i)
	return channels[i]
}
var chan1 chan int
var chan2 chan int
var channels=[]chan int{chan1,chan2}
var num=[]int{1,2,3,4,5}
func test5() {
	

	select{
	case get_chan(0)<-get_number(0):
		fmt.Println("case 1")
	case get_chan(1)<-get_number(1):
		fmt.Println("case 2")
	default:
		fmt.Println("default")
	}
} 
func test6() {
	chan1:=make(chan int,10)
	for i:=0;i<10;i++{
		chan1<-i
	}
	close(chan1)
	synchan:=make(chan struct{})
	fmt.Println("chan1 cap:",cap(chan1))
	fmt.Println("synchan cap:",cap(synchan))

	go func() {
	loop:
		for{
			select{
			case i,ok:=<-chan1:
				if !ok{
					fmt.Println("end")
					break loop
				}
				fmt.Println(i)
			default:
				fmt.Println("default")
			}
		}
		synchan<-struct{}{}
	}()
	<-synchan
}
func test7() {
	chan1:=make(chan int,3)
	synchan:=make(chan struct{})
	go func() {
		for i:=0;i<5;i++{
			chan1<-i
			fmt.Println("send:",time.Now())
			time.Sleep(time.Second)
		}
		close(chan1)
	}()
	go func() {
		loop:
		for{
			time.Sleep(time.Second*2)
			select{
			case i,ok:=<-chan1:
				fmt.Println("receive:",time.Now())
				if !ok{
					fmt.Println("close")
					synchan<-struct{}{}
					break loop
				}else{
					fmt.Println(i)
				}
			}
		}
	}()
	<-synchan
}
func test8() {
	loc, _:= time.LoadLocation("Asia/Chongqing")
    fmt.Println(time.Now().In(loc))

	timer:=time.NewTimer(time.Second*2)
	fmt.Printf("now: %v\n",time.Now().In(loc))
	expirationTime:=<-timer.C
	fmt.Printf("exp: %v\n",expirationTime.In(loc))
	fmt.Printf("stop: %v\n",timer.Stop())
}
func test9() {
	chan1:=make(chan int)
	go func() {
		time.Sleep(time.Second)
		chan1<-1
		close(chan1)
	}()
	timeout:=time.Millisecond*500
	var timer *time.Timer
	for{
		if timer==nil{
			timer=time.NewTimer(timeout)
		}else{
			timer.Reset(timeout)
		}
		select{
			case i,ok:=<-chan1:
				if !ok{
					fmt.Println("end")
					return
				}else{
					fmt.Println(i)
				}
				
			// case <-time.NewTimer(time.Millisecond*500).C:
			case <-timer.C:
				fmt.Println("timeout")
		}
	}
	
	
}
func test10() {
	chan1:=make(chan int)
	ticker:=time.NewTicker(time.Second)
	done:=make(chan struct{})
	// synchan:=make(chan struct{})
	go func() {
		loop:
		for {
			select{
			case <-ticker.C:
				fmt.Println("ticker.C")
				select{
				case chan1<-1:
				case chan1<-2:
				case chan1<-3:
				}
			case <-done:
				break loop
			}
		}
		fmt.Println("end")
		// synchan<-struct{}{}
	}()
	sum:=0
	for i:=range chan1{
		fmt.Println("i:",i)
		sum+=i
		if sum>10{
			fmt.Println(sum)
			ticker.Stop()
			done<-struct{}{}
			// close(ticker.C)
			// if _,ok:=<-ticker.C;!ok{
			// 	fmt.Println("last ticker")
			// }
			break
		}
	}
	fmt.Println("done")
	// <-synchan
	// time.Sleep(time.Second)
}
func test11() {
	defer func() {
		if p:=recover();p!=nil{
			fmt.Printf("%#v\n",p)
		}
	}()
	var mutex sync.Mutex
	mutex.Lock()
	mutex.Unlock()
	fmt.Println("..............")
	mutex.Unlock()
}
func test() {
	var rwm sync.RWMutex
	for i:=0;i<3;i++{
		go func(i int) {
			rwm.RLock()
			fmt.Println("lock:",i)
			time.Sleep(time.Second*2)
			rwm.RUnlock()
			fmt.Println("Unlock:",i)
		}(i)
	}
	time.Sleep(time.Millisecond*300)
	rwm.Lock()
	fmt.Println("done")
}
func main() {
	test()
	time.Sleep(time.Second)
}