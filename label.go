package main
import "fmt"


func main(){
	LABEL1:
	for i:=0;i<=5;i++ {
		for j:=0;j<=5;j++{
			if j==4 || i==4 {
				fmt.Println("")
				continue LABEL1
			}
		fmt.Printf("i is %d , and j is %d\n",i,j)
		}
	}
}
