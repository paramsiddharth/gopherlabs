package student

import "fmt"

type gender int

const (
	Male gender = iota
	Female
	NonBinary
	Unknown
)

type Student struct {
	Name   string
	ID     int32
	Gender gender
}

func (s Student) Display() {
	fmt.Printf(
		`ID: %d
Name: %s
Gender: %s
`,
		s.ID, s.Name, genderEnumToString(s.Gender))
}

func genderEnumToString(g gender) string {
	switch g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	case NonBinary:
		return "Non-Binary"
	case Unknown:
		return "Unknown"
	}
	return ""
}
