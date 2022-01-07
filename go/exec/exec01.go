package exec

import "fmt"

/* 买卖股票的最佳时机 */
var (
	ans int
)

func MaxProfit(prices []int) {
	for i := 1; i < len(prices); i++ {
		ans += clacMax(0, prices[i]-prices[i-1])
	}
	fmt.Println(ans)
}

func clacMax(a, b int) int {
	if b > a {
		return b
	}
	return a
}
