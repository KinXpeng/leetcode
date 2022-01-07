package exec

import (
	"fmt"
)

/****** 旋转数组 ******/

func RotateArr(nums []int, k int) {
	k = k % len(nums)
	copy(nums, append(nums[len(nums)-k:], nums[0:len(nums)-k]...))
	fmt.Println(nums)
}
