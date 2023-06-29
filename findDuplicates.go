package main


import (
		"fmt"
		"bufio"
		"os"
	)


func main(){
	//making a map with key as string (lines enterred ) and value as int as the amount of time the words input as present
	counts:=make(map[string]int)

	//Scanner i presume is used to Scan text from os.Stdin which is equivalent of input() or cin  
	//While I understood the code I think the rabit hole goes way deep than what it seems (I don't know what buffers are SO Here we go)
	// I will enter my findings and the resource i used here or maybe in a blog post , lemme check that shit first
	fmt.Printf("Program Started")
	input:=bufio.NewScanner(os.Stdin)
	for input.Scan(){
		if input.Text()==""{
			break
		}
		counts[input.Text()]++
	}
	for key,value:=range counts {
		if value>1 {
			fmt.Printf("%d\t%s\n",value,key)
		}
	} 
}



//