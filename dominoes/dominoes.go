package dominoes

type Domino [2]int

func MakeChain(input []Domino) ([]Domino, bool) {
	if len(input) == 0 {
		return []Domino{}, true
	}
	if len(input) == 1 {
		if input[0][0] == input[0][1] {
			return input, true
		}
		return nil, false
	}

	used := make([]bool, len(input))
	var result []Domino

	var backtrack func([]Domino) bool
	backtrack = func(curr []Domino) bool {
		if len(curr) == len(input) {
			// Check if the chain can form a loop: last number must match first
			if curr[0][0] == curr[len(curr)-1][1] {
				result = append([]Domino(nil), curr...) // copy
				return true
			}
			return false
		}

		for i, d := range input {
			if used[i] {
				continue
			}

			for _, tile := range []Domino{d, {d[1], d[0]}} {
				if len(curr) == 0 || curr[len(curr)-1][1] == tile[0] {
					used[i] = true
					if backtrack(append(curr, tile)) {
						return true
					}
					used[i] = false
				}
			}
		}
		return false
	}

	if backtrack([]Domino{}) {
		return result, true
	}
	return nil, false
}
