package knapsack

type Item struct {
	Weight, Value int
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maximumWeight
func Knapsack(maximumWeight int, items []Item) int {
	n := len(items)
	if n == 0 || maximumWeight == 0 {
		return 0
	}

	// Create a DP table where dp[i][w] is max value using first i items with capacity w
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, maximumWeight+1)
	}

	for i := 1; i <= n; i++ {
		for w := 0; w <= maximumWeight; w++ {
			if items[i-1].Weight > w {
				// Can't take the current item
				dp[i][w] = dp[i-1][w]
			} else {
				// Max of taking it or not taking it
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-items[i-1].Weight]+items[i-1].Value)
			}
		}
	}

	return dp[n][maximumWeight]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
