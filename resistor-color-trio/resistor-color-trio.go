package resistorcolortrio

import "fmt"

var colorDigits = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

var colorMultipliers = map[string]int{
	"black":  1,
	"brown":  10,
	"red":    100,
	"orange": 1000,
	"yellow": 10000,
	"green":  100000,
	"blue":   1000000,
	"violet": 10000000,
	"grey":   100000000,
	"white":  1000000000,
}

func Label(colors []string) string {
	if len(colors) < 3 {
		return ""
	}

	// Only use the first three colors
	colors = colors[:3]

	d1, ok1 := colorDigits[colors[0]]
	d2, ok2 := colorDigits[colors[1]]
	mult, ok3 := colorMultipliers[colors[2]]

	if !ok1 || !ok2 || !ok3 {
		return ""
	}

	resistance := int64((d1*10 + d2) * mult)

	switch {
	case resistance >= 1_000_000_000:
		return fmt.Sprintf("%d gigaohms", resistance/1_000_000_000)
	case resistance >= 1_000_000:
		return fmt.Sprintf("%d megaohms", resistance/1_000_000)
	case resistance >= 1_000:
		return fmt.Sprintf("%d kiloohms", resistance/1_000)
	default:
		return fmt.Sprintf("%d ohms", resistance)
	}
}
