package main

import ("fmt")

func main(){
	var arr [5]int 
	for i:=0;i<len(arr);i++ {
		arr[i]=i*i
	}

	arr2:=&arr


	arr2[0]=50
	
	if arr2==&arr {
		fmt.Println("TRUE")
	}
	for _,v:=range arr2{
		fmt.Printf("%d\t",v)
	}
	fmt.Println("")
	for _,v:=range arr {
		fmt.Printf("%d\t",v)
	}
}
