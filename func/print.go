package main
import "fmt"

func main() { // main function started
    fmt.Println("In main before calling greeting")

    greeting() // greeting function invoked 

    fmt.Println("In main after calling greeting") // executed after greeting function
} // main function ended

// greeting function declared
func greeting() {
    fmt.Println("In greeting: Hi!!!!!")
}