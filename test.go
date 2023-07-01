//Tests generally results in either result or err
//function can return more than one value and in go 
//you have to recieve either all values or put in the placeholder

package main

import (
	"fmt"
	"strconv"
	"os"
)

func main(){
	var org string="ABC"
	var an int
	var err error

	an,err=strconv.Atoi(org)
	if err!=nil{
		fmt.Printf("orig %s is not an integer -exiting with error status %v\n",an,err)
		os.Exit(1)	
	} else{
		fmt.Printf("The integer produced is %v\n",an)
	}
}