func maxProfit(prices []int) int {
	left, right := 0, 1

	maxProfit := 0

	for right < len(prices) {
		if prices[right] > prices[left] {
			profit := prices[right] - prices[left]

			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			left = right
		}

		right++
	}

	return maxProfit
}
