package main

import (
	"fmt"
	"time"
)


func main(){

	t:=time.Now()
	fmt.Printf("%2d.%2d.%4d\n",t.Day(),t.Month(),t.Year())

}