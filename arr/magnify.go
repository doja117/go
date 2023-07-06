package main

import "fmt"

func main(){
	s:=[]int{1,2,3}

	z:=10

	ans:=make([]int,z)
	//copy(to,from)
	copy(ans,s)

	ans1:=make([]int,2)
	copy(ans1,s)
	for _,v:=range ans1{
		fmt.Printf("%d\t",v)
	}
}