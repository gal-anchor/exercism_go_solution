package allergies

var allergens = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

// Allergies returns a slice of allergen names the person is allergic to.
func Allergies(allergies uint) []string {
	result := []string{}
	for i, allergen := range allergens {
		if allergies&(1<<i) != 0 {
			result = append(result, allergen)
		}
	}
	return result
}

// AllergicTo returns true if the person is allergic to a specific allergen.
func AllergicTo(allergies uint, allergen string) bool {
	for i, a := range allergens {
		if a == allergen {
			return allergies&(1<<i) != 0
		}
	}
	return false
}
