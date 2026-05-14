package pgms

import "fmt"

func stockBuyAndSellV1(prices []int) int {
	n := len(prices)
	profit := 0
	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < n; j++ {
			profit = max(profit, prices[j] - prices[i])
		}
	}

	return profit
}

func stockBuyAndSellV2(prices []int) int {
	n := len(prices)
	profit := 0
	minSoFar := prices[0]

	for i := 1; i < n; i++ {
		minSoFar = min(minSoFar, prices[i])
		profit = max(profit, prices[i] - minSoFar)
	}

	return profit
}

func Program022() {
	prices := []int{7, 10, 1, 3, 6, 9, 2}
	fmt.Println(stockBuyAndSellV2(prices))
}