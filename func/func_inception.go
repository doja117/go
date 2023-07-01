package main
import (
	"fmt"
	"math"	
)

func f1(x,y,z int)(int){
	return x+y+z
}
func f2(x,y int)(int,int,int){
	if y!=0 {
		return int(math.Pow(float64(x),float64(y))),int(math.Pow(float64(y),float64(x))),int(x/y)
	}
	return int(math.Pow(float64(x),float64(y))),int(math.Pow(float64(y),float64(x))),0
}

func main() {
	fmt.Printf("%d\n",f1(f2(5,7)))
}

