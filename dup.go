package main

import (
	"fmt"
	"os"
	"bufio"
	
)



func main(){
	counts:=make(map[string]int)
	files:=os.Args[1:]
	if(len(files)==0){
		countLines(os.Stdin,counts)
	}	else{
		for _,arg:=range files{
			f,err:=os.Open(arg)
			if err!=nil{
				fmt.Fprintf(os.Stderr,"dup2 :%v\n",err)
				continue
			}	
			countLines(f,counts)
			f.Close()
		}
	}
	for line,v:=range counts{
		if v>1{
			fmt.Printf("%d\t%s\n",v,line)
		}
	}

}

func countLines(f *os.File,counts map[string]int){
	input:=bufio.NewScanner(f)
	for input.Scan(){
		if input.Text()!=""{
			counts[input.Text()]++
		}
		break
		
	}
}