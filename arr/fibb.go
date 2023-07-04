package main

import (
	"fmt"
	"os"
)


func main(){

	var arr []int
	arr=append(arr,0)
	arr=append(arr,1)


	z:=int(os.Args[1])
	
	
	for i:=2;i<z;i++ {
		arr=append(arr,arr[i-1]+arr[i-2])
	}

	for _,v:=range arr {
		fmt.Printf("%d\t",v)
	}
}
