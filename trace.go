package main

import (
	"fmt"
)


func trace(s string){
	fmt.Printf("entering : ",s)
}
func untrace(s string){
	fmt.Printf("exiting : ",s)
}

func a(){
	trace("a")
	defer untrace("a")
	fmt.Printf("in a \n")
}

func b(){
	trace("b")
	defer untrace("b")
	fmt.Printf("in b")
	a()
	fmt.Printf("\n")
}


func main(){
	b()
}