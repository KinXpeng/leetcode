package exec

import (
	"fmt"
	"math/rand"
	"time"
)

/* 生成随机数组并翻转数组 */
func FlipArray() {
	var (
		randArr [7]int
		temp    = 0 // 临时变量，存储中间数据
	)
	num := len(randArr)
	// 生成1-100的随机数组
	rand.Seed(time.Now().Unix())
	for i := 0; i < num; i++ {
		randArr[i] = rand.Intn(100)
	}
	fmt.Printf("生成的随机数组为%v\n", randArr)
	// 翻转数组
	for j := 0; j < num/2; j++ {
		temp = randArr[num-1-j]
		randArr[num-1-j] = randArr[j]
		randArr[j] = temp
	}
	fmt.Printf("翻转后的数组为%v\n", randArr)
}
