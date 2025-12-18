package change

import (
	"errors"
	"sort"
)

func Change(coins []int, target int) ([]int, error) {
	if target == 0 {
		return []int{}, nil
	}
	if target < 0 {
		return nil, errors.New("target amount cannot be negative")
	}

	// dp[i] will store the minimum number of coins to reach amount i
	dp := make([]int, target+1)
	// parent[i] stores which coin was added to reach amount i
	parent := make([]int, target+1)

	// Initialize dp with a value larger than any possible solution
	for i := 1; i <= target; i++ {
		dp[i] = target + 1
	}

	// Build the DP table
	for _, coin := range coins {
		for i := coin; i <= target; i++ {
			if dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
				parent[i] = coin
			}
		}
	}

	// If the target is unreachable
	if dp[target] > target {
		return nil, errors.New("no combination of coins can satisfy the target")
	}

	// Backtrack to find the coins used
	result := []int{}
	curr := target
	for curr > 0 {
		coin := parent[curr]
		result = append(result, coin)
		curr -= coin
	}

	// Return sorted result for consistency
	sort.Ints(result)
	return result, nil
}
