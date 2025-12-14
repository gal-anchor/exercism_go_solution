package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer < 1 {
		return len(layers) * 2
	}
	return len(layers) * timePerLayer
}

func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0.0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendList []string, myList []string) {
	(myList)[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	needs := make([]float64, len(amounts))

	for index, amount := range amounts {
		needs[index] = amount * float64(portions) / 2
	}

	return needs
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
