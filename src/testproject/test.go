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
func test() {
 	go fmt.Println("wo")
} 
func main() {
	test()
	time.Sleep(time.Second)
}