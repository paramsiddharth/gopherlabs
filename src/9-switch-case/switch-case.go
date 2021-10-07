package main

func main() {
	truth := 9 == 3*3

	// Boolean switch-case
	switch {
	case truth:
		println("True!")
	default:
		println("False!")
	}

	x := "Param"

	// Value-based
	switch x {
	case "Henry":
		println("Henry's here!")
	default:
		println("Someone's here!")
	}

	y := 3

	// Multiple
	switch y {
	case 1, 3, 5:
		println("Odd~!")
	default:
		println("Event~!")
	}
}
