package railfence

func Encode(message string, rails int) string {
	if rails <= 1 || rails >= len(message) {
		return message
	}

	// Create a slice of strings for each rail
	rows := make([][]rune, rails)
	rail := 0
	dir := 1 // 1 for down, -1 for up

	for _, ch := range message {
		rows[rail] = append(rows[rail], ch)
		rail += dir
		if rail == 0 || rail == rails-1 {
			dir *= -1
		}
	}

	// Concatenate all rows
	result := ""
	for _, row := range rows {
		result += string(row)
	}
	return result
}

func Decode(message string, rails int) string {
	if rails <= 1 || rails >= len(message) {
		return message
	}

	// Step 1: Determine the zigzag pattern positions
	n := len(message)
	pos := make([]int, n)
	rail := 0
	dir := 1
	for i := 0; i < n; i++ {
		pos[i] = rail
		rail += dir
		if rail == 0 || rail == rails-1 {
			dir *= -1
		}
	}

	// Step 2: Count letters per rail
	counts := make([]int, rails)
	for _, p := range pos {
		counts[p]++
	}

	// Step 3: Fill rails with letters
	railsData := make([][]rune, rails)
	idx := 0
	for r := 0; r < rails; r++ {
		railsData[r] = []rune(message[idx : idx+counts[r]])
		idx += counts[r]
	}

	// Step 4: Reconstruct the message
	result := make([]rune, n)
	railIndices := make([]int, rails)
	for i := 0; i < n; i++ {
		r := pos[i]
		result[i] = railsData[r][railIndices[r]]
		railIndices[r]++
	}

	return string(result)
}
