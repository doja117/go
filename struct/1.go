package main

import "fmt"
import "reflect"

type struct1 struct {
i1 int
f1 float64
s1 string
}

func main(){
	//new function returns pointer 

	ms:=new(struct1)

	ms.i1=10
	ms.f1=10.123
	ms.s1="Hello World"


	fmt.Printf("%s\t",reflect.TypeOf(*ms))
	

}
