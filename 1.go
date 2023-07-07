package main

import "fmt"

func returnSorted(arr []int) []int {
	if len(arr) == 1 || len(arr) == 0 {
		return arr
	}

	ans := make([]int, 0)

LOOP:
	z := 1000000000
	it := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] < z {
			z = arr[i]
			it = i
		}
	}
	ans = append(ans, z)
	if len(arr) == 1 {
		ans = append(ans, arr[0])
		return ans
	} else {
		if it == len(arr)-1 {
			arr = arr[:it]
			goto LOOP
		} else {
			arr = append(arr[:it], arr[it+1:]...)
			goto LOOP
		}
	}
	return ans

}
func main() {

	arr := []int{5, 3, 2, 1, 4, 4}
	fmt.Println(returnSorted(arr[:len(arr)-1]))
}
