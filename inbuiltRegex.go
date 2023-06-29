package main

import (
	"fmt"
	"strings"
)

func main(){
	var str string="Hi, I am Doja."
	fmt.Printf("The position of the first / Doja / is ")
	fmt.Printf("%d\n",strings.Index(str,"Doja"))
	fmt.Printf("The position of the first / Hi / is ")
	fmt.Printf("%d\n",strings.Index(str,"Hi"))

	oldName:="Doja"
	newName:="Master Chief"
	newStr:=strings.Replace(str,oldName,newName,-1)

	// A new copy is returned while the old remains intact , maybe because of strings being immutable 
	fmt.Println(newStr)
}	
