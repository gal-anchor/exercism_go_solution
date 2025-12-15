package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type TeamStats struct {
	Wins   int
	Draws  int
	Losses int
}

func (ts TeamStats) Matches() int {
	return ts.Wins + ts.Draws + ts.Losses
}

func (ts TeamStats) Points() int {
	return ts.Wins*3 + ts.Draws
}

// Tally reads tournament results from reader, computes scores, and writes a formatted table to writer.
func Tally(reader io.Reader, writer io.Writer) error {
	stats := make(map[string]*TeamStats)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue // skip empty lines or comments
		}

		parts := strings.Split(line, ";")
		if len(parts) != 3 {
			return fmt.Errorf("invalid line format: %s", line)
		}

		team1, team2, result := parts[0], parts[1], parts[2]

		if _, exists := stats[team1]; !exists {
			stats[team1] = &TeamStats{}
		}
		if _, exists := stats[team2]; !exists {
			stats[team2] = &TeamStats{}
		}

		switch result {
		case "win":
			stats[team1].Wins++
			stats[team2].Losses++
		case "loss":
			stats[team1].Losses++
			stats[team2].Wins++
		case "draw":
			stats[team1].Draws++
			stats[team2].Draws++
		default:
			return fmt.Errorf("invalid result: %s", result)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Prepare teams sorted by points descending, then name ascending
	teams := make([]string, 0, len(stats))
	for team := range stats {
		teams = append(teams, team)
	}

	sort.Slice(teams, func(i, j int) bool {
		a, b := stats[teams[i]], stats[teams[j]]
		if a.Points() != b.Points() {
			return a.Points() > b.Points()
		}
		return teams[i] < teams[j]
	})

	// Write table header
	fmt.Fprintf(writer, "%-31s| MP |  W |  D |  L |  P\n", "Team")

	// Write each team
	for _, team := range teams {
		s := stats[team]
		fmt.Fprintf(writer, "%-31s| %2d | %2d | %2d | %2d | %2d\n",
			team, s.Matches(), s.Wins, s.Draws, s.Losses, s.Points())
	}

	return nil
}
