package main

import "fmt"
const (
	WIDTH=1920
	HEIGHT=1080
)
type pixel int
var screen [WIDTH][HEIGHT]pixel
func main(){
	values:=[][]int{}
	row1:=[]int{1,2,3}
	row2:=[]int{4,5,6}

	values=append(values,row1)
	values=append(values,row2)

	for _,v:=range values{
		for _,v1:=range v {
			fmt.Printf("%d\t",v1)
		}
	fmt.Printf("\n")
	}
	for y:=0;y<HEIGHT;y++{
		for x:=0;x<WIDTH;x++ {
			screen[x][y]=0
		}
	}
}
