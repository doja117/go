package main
//youtube.com/watch?v=ENxfg9rS5dc
import (
	"fmt"
	"bytes"
	)


func main(){
	buf:=new(bytes.Buffer)
	buf.Write([]byte("Foo"))
	buf.WriteString("Jerwiqpoje")
	fmt.Println(buf.Len())
	fmt.Println(buf.String())
}
