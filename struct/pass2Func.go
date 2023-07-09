package main

import (
	"fmt"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func returnNameInCaps(p *Person) string{
	return strings.ToUpper(p.fname)+" "+strings.ToUpper(p.lname)

}
func main(){
	p:=new(Person)
	p.fname="sid"

	p.lname="oj"
	fmt.Println(returnNameInCaps(p))
}
