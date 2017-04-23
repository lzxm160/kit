package main
import (
	"fmt"
	"time"
)
func  test() {
 	names:=[]string{"1","2","3","4","5"}
 	for _,name:=range names{
 		go func(who string) {
 			fmt.Println(who)	
 		}(name)
 	}
 } 

func main() {
	test()
	time.Sleep(time.Second)
}