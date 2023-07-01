package main

import (
	"fmt"
)

func main(){
	str:="Go is a decent language"
	for pos,char:=range str{
		fmt.Printf("%v:%c\n",pos,char)
	}	
}