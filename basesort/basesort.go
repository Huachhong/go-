package basesort

import "fmt"

func BubbleSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if arr[i] > arr[j] {
				//变量交换 第一种方法
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp

				/*				//变量交换 第二种方法
								arr[i] = arr[i] + arr[j]
								arr[j] = arr[i] - arr[j]
								arr[i] = arr[i] - arr[j]*/

				/*				//变量交换 第三种方法
								arr[i], arr[j] = arr[j], arr[i]*/
			}
		}
	}
	return arr
}

func TestSort() {
	arr := []int{10, 8, 2, 1, 0, 7, 3, 6, 5, 9, 4}
	var res []int
	res = BubbleSort(arr)
	fmt.Printf("\n sort arr = %v", res)
}
