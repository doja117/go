package main

import "fmt"



func main(){

	var arr []int
	var z int
	for i:=0;i<1e9;i++ {
		if z!=cap(arr){
		fmt.Printf("Length: %d Capacity: %d\n",len(arr),cap(arr))
		z=cap(arr)
		}
		arr=append(arr,i)
	}

}
