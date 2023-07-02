package main

import (
	"fmt"
	"time"
)

func printLoop(x,y int){
	tm:=time.Now()
	time.Sleep(2*time.Second)
	fmt.Print("%s", tm)
	for i:=x;i<y;i++ {
		fmt.Printf("%d\t",i)
	}

	tm1:=time.Now()
	fmt.Print("%s",tm1)
	fmt.Println("")
}

func main(){
	printLoop(0,5)
	defer printLoop(5,10)
	printLoop(10,15)
}