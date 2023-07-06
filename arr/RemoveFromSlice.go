package main

import "fmt"

func removeSlice(sl []int, index int) []int {
	sl[index] = sl[len(sl)-1]
	return sl[:len(sl)-1]
}
func removeSliceWithOrder(sl []int, index int) []int {
	return append(sl[:index], sl[index+1:]...)
}
func main() {
	arr := []int{1, 2, 3, 5, 6, 7}
	fmt.Println(removeSlice(arr, 1))
	fmt.Println(removeSliceWithOrder(arr, 1))
}
