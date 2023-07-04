package main

import "fmt"


func main(){
	var p *[]int=new([]int)

	for i:=0;i<10;i++ {
		*p=append(*p,i)
	}
	
	for _,v:=range *p {
		fmt.Printf("%d\t",v)
	}
	
}
