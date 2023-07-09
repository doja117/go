package main

import "fmt"

type R struct {
	z float64
	int
	string
}

func main() {
	p := R{3.14, 3, "Saurabh"}
	fmt.Printf("%.2f\n", p.z)
	fmt.Printf("%s", p.string)
}
