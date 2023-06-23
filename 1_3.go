package main

import (
	"fmt"
	"time"
	"os"
	"strings"
	)


func stupidMethod(){
	var s,sep string
	t1:=time.Now()
	for i:=0;i<len(os.Args);i++ {
		s+=os.Args[i]+sep
		sep=" "
	}
	//fmt.Printf(s)
	t2:=time.Now()

	t_stup:=t2.Sub(t1)
	fmt.Println("%d",t_stup)
}

func standardMethod(){
	t1:=time.Now()
	strings.Join(os.Args[1:]," ")
	t2:=time.Now()

	t_stnd:=t2.Sub(t1)
	fmt.Println("%d",t_stnd)
}
func main(){
	

	fmt.Printf("Time for stupid Method")
	stupidMethod()


	fmt.Printf("Time for standard Method")
	standardMethod()

}


