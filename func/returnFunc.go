package main

import "fmt"
/*
func getInc(b int) func(x int) int {
	return func(x int){
	return b+x}	
}
*/


func main(){
	p2:=Add2()
	fmt.Printf("%d\t",p2(2))
	TwoAdder:=Adder(2)
	fmt.Printf("The result is : %v\n",TwoAdder(2))
}


func Add2(func (b int) int){
	return func(b int) int{
		return b+2
	}
}

func Adder(a int) (func (b int) int){
	return func(b int) int {
		return a+b
	}
}
