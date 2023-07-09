package main

import "fmt"

func main() {
	m := make(map[int]string)
	for i := 1; i < 101; i++ {
		if i%15 == 0 {
			m[i] = "FIZZBUZZ"
		} else if i%5 == 0 {
			m[i] = "FIZZ"
		} else if i%3 == 0 {
			m[i] = "BUZZ"
		} else {
			m[i] = string(i)
		}
	}

	for k, v := range m {
		fmt.Printf("%d:\t%s\n", k, v)
	}

}
