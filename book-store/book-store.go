package bookstore

import "fmt"

func Cost(books []int) int {
	counts := make([]int, 5)
	for _, b := range books {
		if b >= 1 && b <= 5 {
			counts[b-1]++
		}
	}

	cache := map[string]int{}

	var dp func([]int) int
	dp = func(c []int) int {
		key := fmt.Sprintf("%v", c)
		if val, ok := cache[key]; ok {
			return val
		}

		// Base case: all counts zero
		allZero := true
		for _, x := range c {
			if x > 0 {
				allZero = false
				break
			}
		}
		if allZero {
			return 0
		}

		best := int(^uint(0) >> 1)

		// Generate all non-empty subsets of books that can form a group
		for mask := 1; mask < 32; mask++ { // 5 books → 2^5 subsets
			valid := true
			groupSize := 0
			next := make([]int, 5)
			copy(next, c)
			for i := 0; i < 5; i++ {
				if mask&(1<<i) != 0 {
					if c[i] == 0 {
						valid = false
						break
					}
					next[i]--
					groupSize++
				}
			}
			if !valid || groupSize == 0 {
				continue
			}

			price := groupSize * 800
			switch groupSize {
			case 2:
				price = price * 95 / 100
			case 3:
				price = price * 90 / 100
			case 4:
				price = price * 80 / 100
			case 5:
				price = price * 75 / 100
			}

			cost := price + dp(next)
			if cost < best {
				best = cost
			}
		}

		cache[key] = best
		return best
	}

	return dp(counts)
}
