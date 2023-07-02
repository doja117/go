package main

//Testing the classic and wrong way of timing how efficient a programing language is , Python devs mostly hate this . Impractical but doing it for the record
//Time more taking in printing rather than performing logic
import (
"fmt"
"time"
)


func main(){
t0:=time.Now()
for i:=0;i<1e6;i++{
fv:=func(i int){fmt.Printf("%d\t",i)}
if((i%23)==0) {
fv(i)
}
}
fmt.Printf("\n")
elapsed:=time.Since(t0)
fmt.Printf("page took %s seconds",elapsed)
time.Sleep(time.Second)

}
