package main

import (
	"fmt"
	"strconv"
)

func fibb(x int32)int32{ 
	switch x{
	case 0:return 0
	case 1:return 1
	default:return fibb(x-1)+fibb(x-2)
}
}



func main(){
	
	var y int32
	for y=1;y<15;y++ {
	fmt.Printf("%d\n",fibb(y))
	}
	
}