package main

import (
	"fmt"
	
)

func passTheDouchie(length *float64,count *int){
	if *length<=0 {
		return 
	}
	number:=*length/10
	*length-=number
	*count+=1
}

func main(){
	var length float64=10.00
	t:=&length
	count:=0
	i:=1
	for length>0 {
		passTheDouchie(t,&count)
		fmt.Printf("%d : %.2f\n",i,*t)
		if *t<=0.01 {break}
		//time.Sleep(1*time.Second)
		i++
	}
	fmt.Printf("Number of passes : %d\n",count)
}