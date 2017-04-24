package main
import (
	"fmt"
	"time"
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
func test() {
	countmap:=make(chan map[string]*Counter,1)
	sync2:=make(chan struct{},2)
	go receive(countmap,sync2)
	go sends(countmap,sync2)
	
	<-sync2

	<-sync2
} 
func main() {
	test()
	time.Sleep(time.Second)
}