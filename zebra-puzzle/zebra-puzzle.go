package zebra

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

func SolvePuzzle() Solution {
	nations := []string{"Englishman", "Spaniard", "Ukrainian", "Norwegian", "Japanese"}
	perms := generatePermutations([]int{0, 1, 2, 3, 4})

	for _, col := range perms {
		if !checkColors(col) {
			continue
		}
		for _, nat := range perms {
			if !checkNations(nat, col) {
				continue
			}
			for _, bev := range perms {
				if !checkBeverages(bev, nat, col) {
					continue
				}
				for _, smo := range perms {
					if !checkSmokes(smo, nat, bev, col) {
						continue
					}
					for _, pet := range perms {
						if checkPets(pet, nat, smo) {
							// Index 0: Water, Index 4: Zebra
							// We find which house (index 0-4) has that attribute,
							// then see which nation is at that house index.
							waterHouse := findInList(bev, 0)
							zebraHouse := findInList(pet, 4)

							return Solution{
								DrinksWater: nations[nat[waterHouse]],
								OwnsZebra:   nations[nat[zebraHouse]],
							}
						}
					}
				}
			}
		}
	}
	return Solution{}
}

// --- Constraints ---

func checkColors(p []int) bool {
	// 0: Red, 1: Green, 2: Ivory, 3: Yellow, 4: Blue
	// Green (1) is immediately right of Ivory (2)
	return findInList(p, 1) == findInList(p, 2)+1
}

func checkNations(p []int, col []int) bool {
	// 0: Englishman, 1: Spaniard, 2: Ukrainian, 3: Norwegian, 4: Japanese
	// Norwegian (3) in first house (index 0)
	// Englishman (0) in Red (0) house
	// Norwegian (3) next to Blue (4) house
	return p[0] == 3 &&
		findInList(p, 0) == findInList(col, 0) &&
		isNeighbor(findInList(p, 3), findInList(col, 4))
}

func checkBeverages(p []int, nat []int, col []int) bool {
	// 0: Water, 1: Coffee, 2: Tea, 3: Milk, 4: Orange Juice
	// Coffee (1) in Green (1) house
	// Ukrainian (2) drinks Tea (2)
	// Milk (3) in middle house (index 2)
	return findInList(p, 1) == findInList(col, 1) &&
		nat[findInList(p, 2)] == 2 &&
		p[2] == 3
}

func checkSmokes(p []int, nat []int, bev []int, col []int) bool {
	// 0: Old Gold, 1: Kools, 2: Chesterfields, 3: Lucky Strike, 4: Parliaments
	// Kools (1) in Yellow (3) house
	// Lucky Strike (3) drinks Orange Juice (4)
	// Japanese (4) smokes Parliaments (4)
	return findInList(p, 1) == findInList(col, 3) &&
		findInList(p, 3) == findInList(bev, 4) &&
		nat[findInList(p, 4)] == 4
}

func checkPets(p []int, nat []int, smo []int) bool {
	// 0: Dog, 1: Snails, 2: Fox, 3: Horse, 4: Zebra
	// Spaniard (1) has Dog (0)
	// Old Gold (0) has Snails (1)
	// Chesterfields (2) next to Fox (2)
	// Kools (1) next to Horse (3)
	return nat[findInList(p, 0)] == 1 &&
		findInList(p, 1) == findInList(smo, 0) &&
		isNeighbor(findInList(smo, 2), findInList(p, 2)) &&
		isNeighbor(findInList(smo, 1), findInList(p, 3))
}

// --- Helpers ---

func isNeighbor(a, b int) bool {
	return a-b == 1 || b-a == 1
}

func findInList(list []int, val int) int {
	for i, v := range list {
		if v == val {
			return i
		}
	}
	return -1
}

func generatePermutations(arr []int) [][]int {
	var res [][]int
	var helper func([]int, int)
	helper = func(a []int, n int) {
		if n == 1 {
			tmp := make([]int, len(a))
			copy(tmp, a)
			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			helper(a, n-1)
			if n%2 == 0 {
				a[i], a[n-1] = a[n-1], a[i]
			} else {
				a[0], a[n-1] = a[n-1], a[0]
			}
		}
	}
	helper(arr, len(arr))
	return res
}
