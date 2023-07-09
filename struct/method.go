package main
import "fmt"

type twoInt struct {
	x1 int
	x2 int
}

func addTwoInt(z *twoInt) int {
	return z.x1+z.x2
}

func main(){
	z:=&twoInt{1,2}
	fmt.Printf("%d\t",addTwoInt(z))
}
