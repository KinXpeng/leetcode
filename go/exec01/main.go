package main

import "fmt"

// 买卖股票的最佳时机

func main() {
	arr := []int{7, 2, 3, 1, 5, 6}
	m := maxProfit(arr)
	fmt.Println(m)
}

func maxProfit(prices []int) (ans int) {
	for i := 1; i < len(prices); i++ {
		ans += clacMax(0, prices[i]-prices[i-1])
	}
	return ans
}

func clacMax(a, b int) int {
	if b > a {
		return b
	}
	return a
}
