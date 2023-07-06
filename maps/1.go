package main

import "fmt"

func main() {
	myMap := make(map[string]string)

	myMap["Saurabh"] = "Ojha"
	myMap["Harsh"] = "Raj"
	myMap["Go"] = "Lang"

	for k, v := range myMap {
		fmt.Printf("%s\t%s\n", k, v)
	}

	mf := map[int]func() int{ // key type int, and value type func()int
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf)
}
