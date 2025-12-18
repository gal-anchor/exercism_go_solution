package school

import "sort"

// Grade represents a grade level and its students
type Grade struct {
	Level    int
	Students []string
}

// School represents the whole school with students in different grades
type School struct {
	roster map[int][]string
}

// New creates and returns a new empty School
func New() *School {
	return &School{
		roster: make(map[int][]string),
	}
}

// Add adds a student to the given grade
func (s *School) Add(student string, grade int) {
	s.roster[grade] = append(s.roster[grade], student)
	sort.Strings(s.roster[grade]) // Keep students sorted
}

// Grade returns the list of students in a given grade
func (s *School) Grade(level int) []string {
	students := s.roster[level]
	// Return a copy to avoid external modification
	result := make([]string, len(students))
	copy(result, students)
	return result
}

// Enrollment returns a sorted list of all grades with their students
func (s *School) Enrollment() []Grade {
	var grades []Grade
	for level, students := range s.roster {
		// Copy students to avoid external modification
		studentsCopy := make([]string, len(students))
		copy(studentsCopy, students)
		grades = append(grades, Grade{
			Level:    level,
			Students: studentsCopy,
		})
	}
	// Sort grades by Level
	sort.Slice(grades, func(i, j int) bool {
		return grades[i].Level < grades[j].Level
	})
	return grades
}
