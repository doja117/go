package main


import (
	"fmt"
	"time"
)

func increment(x *int,y int){
	*x+=y
}


// Passing by reference is far more better memorywise for bigger data structures like say 
// a struct or an interface
func main(){
	z:=10
	t:=&z
	var k **int=&t
	//A pointer to a pointer
	**k+=10
	increment(t,5)
	fmt.Printf("%d\n",z)

	//Testing nil pointers

	var k1 *int =nil
	var k2 *int =nil
	
	for k1==k2 {
		fmt.Printf("Uninitialize\n")
		time.Sleep(1*time.Second)
		*k1=10
		*k2=15
	}
}
