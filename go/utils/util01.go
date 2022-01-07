package utils

import "fmt"

/*
	@title 冒泡排序
	@params list 需要排序的数组
*/
func BubbleSort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
			}
		}
	}
	fmt.Println(list)
}

/*
	@title 二分法查找(有序数组)
	@params list 需要查找的数组
					leftIndex 查找的起始下标
					rightIndex 查找的结束下标
					findNum 需要查找的元素
*/
func BinaryFind(list []int, leftIndex, rightIndex, findNum int) {
	middle := (leftIndex + rightIndex) / 2
	if leftIndex > rightIndex {
		fmt.Println("Not Found")
		return
	}
	if findNum < list[middle] {
		BinaryFind(list, leftIndex, middle-1, findNum)
	} else if findNum > list[middle] {
		BinaryFind(list, middle+1, rightIndex, findNum)
	} else if findNum == list[middle] {
		fmt.Printf("Found it,index=%v", middle)
	} else {
		fmt.Println("Not found")
	}
}
