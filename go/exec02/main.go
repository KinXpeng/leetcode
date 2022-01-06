package main

import (
	"fmt"
	_	"math"
	_ "unsafe"
)

/****** 旋转数组 ******/ 

func main(){
	arr := []int{1,2,3,4,5,6,7}
	numsArr := rotateArr(arr,3)
	fmt.Println(numsArr)
}

func rotateArr(nums []int,k int) (res []int){
	k = k % len(nums)
	copy(nums, append(nums[len(nums)-k:], nums[0:len(nums)-k]...))
	return nums
}
