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
func test() {
	countmap:=make(chan map[string]int,1)
	sync2:=make(chan struct{},2)
	go func() {
		for{
			if elem,ok:=<-countmap;ok{
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
		// content[""]
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
func main() {
	test()
	time.Sleep(time.Second)
}