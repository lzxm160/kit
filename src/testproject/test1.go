package main
import (
	"fmt"
	"time"
)
func test1() {
 	names:=[]string{"1","2","3","4","5"}
 	for _,name:=range names{
 		go func(who string) {
 			fmt.Println(who)	
 		}(name)
 	}
} 
func test2() {
 	go fmt.Println("wo")
} 
func test3() {
	content:=make(chan string,3)
	sync1:=make(chan struct{},1)
	sync2:=make(chan struct{},2)
	go func() {
		sync1<-struct{}{}
		for{
			if elem,ok:=<-content;ok{
				fmt.Println("sync1:",elem)
				time.Sleep(time.Second)
			}else{
				fmt.Println("sync1 break")
				break
			}
		}
		defer func(){sync2<-struct{}{}}()
	}()
	go func() {
		<-sync1
		fmt.Println("sync2")
		for{
			if elem,ok:=<-content;ok{
				fmt.Println("sync2:",elem)
			}else{
				fmt.Println("sync2 break")
				break
			}
		}
		defer func(){sync2<-struct{}{}}()
	}()
	content<-"a"
	content<-"b"
	content<-"c"
	close(content)
	
	<-sync2

	<-sync2
} 
func main() {
	test()
	time.Sleep(time.Second)
}