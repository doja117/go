package main

import (
	"fmt"
	"os"
	"reflect"
)

func main(){
	f,err:=os.Open("./text.txt")
	if err!=nil{
		os.Exit(1)
	}	else{
		fmt.Printf("jqepwe j\n")
		fmt.Printf("%v",reflect.TypeOf(f))
	}	

}