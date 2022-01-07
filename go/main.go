package main

import (
	utils "exec/utils" //
)

func main() {
	//exec.MaxProfit([]int{7, 2, 3, 1, 5, 6})       // 买卖股票的最佳时机
	//exec.RotateArr([]int{1, 2, 3, 4, 5, 6, 7}, 3) // 旋转数组
	//exec.FlipArray()                              // 生成随机数组并翻转数组

	//utils.BubbleSort([]int{24, 69, 80, 57, 13, 31})           // 冒泡排序
	utils.BinaryFind([]int{13, 24, 31, 57, 69, 80}, 0, 6, 57) // 二分法查找
}
