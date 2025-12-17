package rectangles

func Count(diagram []string) int {
	if len(diagram) == 0 {
		return 0
	}

	pluses := map[[2]int]bool{}
	for y, row := range diagram {
		for x, c := range row {
			if c == '+' {
				pluses[[2]int{x, y}] = true
			}
		}
	}

	count := 0
	for top := range pluses {
		x1, y1 := top[0], top[1]
		for bottom := range pluses {
			x2, y2 := bottom[0], bottom[1]
			if x1 >= x2 || y1 >= y2 {
				continue
			}

			// Ensure other two corners exist
			if !pluses[[2]int{x2, y1}] || !pluses[[2]int{x1, y2}] {
				continue
			}

			// Check top and bottom edges
			valid := true
			for x := x1 + 1; x < x2; x++ {
				if !isHorizontalEdge(diagram[y1], x) || !isHorizontalEdge(diagram[y2], x) {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}

			// Check left and right edges
			for y := y1 + 1; y < y2; y++ {
				if !isVerticalEdge(diagram, x1, y) || !isVerticalEdge(diagram, x2, y) {
					valid = false
					break
				}
			}

			if valid {
				count++
			}
		}
	}

	return count
}

func isHorizontalEdge(row string, x int) bool {
	if x >= len(row) {
		return false
	}
	return row[x] == '-' || row[x] == '+'
}

func isVerticalEdge(diagram []string, x, y int) bool {
	if y >= len(diagram) || x >= len(diagram[y]) {
		return false
	}
	return diagram[y][x] == '|' || diagram[y][x] == '+'
}
