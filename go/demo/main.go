package main

import "fmt"

func main() {
	list := []int{24, 69, 80, 57, 13, 31}
	result := sort(list)
	fmt.Println(result)
}

/* 实现冒泡排序 */
func sort(list []int) []int {
	length := len(list)
	num := 0
	for num < length-1 {
		for i := 0; i < length-1-num; i++ {
			if list[i] > list[i+1] {
				temp := list[i+1]
				list[i+1] = list[i]
				list[i] = temp
			}
		}
		num++
	}
	return list
}
