package yacht

import "sort"

func Score(dice []int, category string) int {
	switch category {
	case "ones":
		return sumDice(dice, 1)
	case "twos":
		return sumDice(dice, 2)
	case "threes":
		return sumDice(dice, 3)
	case "fours":
		return sumDice(dice, 4)
	case "fives":
		return sumDice(dice, 5)
	case "sixes":
		return sumDice(dice, 6)
	case "full house":
		return fullHouse(dice)
	case "four of a kind":
		return fourOfAKind(dice)
	case "little straight":
		if isStraight(dice, 1, 5) {
			return 30
		}
	case "big straight":
		if isStraight(dice, 2, 6) {
			return 30
		}
	case "choice":
		return sumAll(dice)
	case "yacht":
		if allSame(dice) {
			return 50
		}
	}
	return 0
}

// Helper functions
func sumDice(dice []int, target int) int {
	sum := 0
	for _, d := range dice {
		if d == target {
			sum += d
		}
	}
	return sum
}

func sumAll(dice []int) int {
	sum := 0
	for _, d := range dice {
		sum += d
	}
	return sum
}

func allSame(dice []int) bool {
	for _, d := range dice {
		if d != dice[0] {
			return false
		}
	}
	return true
}

func countDice(dice []int) map[int]int {
	count := make(map[int]int)
	for _, d := range dice {
		count[d]++
	}
	return count
}

func fullHouse(dice []int) int {
	count := countDice(dice)
	if len(count) == 2 {
		for _, v := range count {
			if v != 2 && v != 3 {
				return 0
			}
		}
		return sumAll(dice)
	}
	return 0
}

func fourOfAKind(dice []int) int {
	count := countDice(dice)
	for num, c := range count {
		if c >= 4 {
			return num * 4
		}
	}
	return 0
}

func isStraight(dice []int, start, end int) bool {
	sort.Ints(dice)
	for i := start; i <= end; i++ {
		if dice[i-start] != i {
			return false
		}
	}
	return true
}
