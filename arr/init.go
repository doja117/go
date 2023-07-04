package main

import "fmt"

func main(){
	s1:="Saurabh"

	s2:=make([] string ,10)

	for i:=0;i<len(s2);i++ {
		s2[i]=s1
	}

	for _,v:=range s2 {
		for _,v1:=range v {
			fmt.Printf("%c\t",v1)
		}
	fmt.Println("")
	}
}
