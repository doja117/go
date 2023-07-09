package main
import "fmt"

type List []int

func (l *List) Len() int {return len(*l)}

func (l *List) Append(val int){return append(*l,val)}
func main(){
	var lst List
	lst.Append(1)
	lst.Append(2)
	fmt.Printf("%d\n",lst.Len())
}
