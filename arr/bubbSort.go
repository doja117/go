package main

import "fmt"

//https://www.youtube.com/watch?v=AL_C9nF_0ss&ab_channel=AnthonyGG

func bubblSort(arr []int) []int {
	if len(arr) == 1 || len(arr) == 0 {
		return arr
	}
	count := 0

LOOP:
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] && i+2 != len(arr) {
			//fmt.Println(arr)
			//fmt.Printf("%d\t%d\t%d\t", arr[i], arr[i+1], count)
			count += 1

			new_arr := append(arr[:i], arr[i+1], arr[i])
			new_arr = append(new_arr, arr[i+2:]...)
			arr = new_arr
			//fmt.Println(arr)
			goto LOOP
		} else if arr[i] > arr[i+1] && i+2 == len(arr) {
			new_arr := append(arr[:i], arr[i+1], arr[i])
			arr = new_arr
			goto LOOP
		}
	}
	return arr
}
func main() {
	arr := []int{9, 6, 7, 4, 1}
	fmt.Println(bubblSort(arr))
}
