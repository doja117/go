package main

import (
	"fmt"
)

func insertSlice(slice,insertionString []string,index int) []string {
	new_s:=slice[index:]
	old_s:=slice[0:index-1]
	ans:=append(insertionString,new_s)
	ans:=append(old_s,ans)
	return ans
}
func main(){
	s1:="MNOPQR"
	s2:="ABC"
	ans:=insertSlice(s1,s2,0)

	for _,v:=range ans {
		fmt.Printf("%d",v)
	}
}