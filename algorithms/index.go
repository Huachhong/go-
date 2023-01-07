package algorithms

import "fmt"

func Test() {
	arr := []int{8, 0, 3, 7, 2}
	sortarr := mergeSort(arr)
	fmt.Println(sortarr)
}
