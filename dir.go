package main

type Dir byte

const (
	DIR_RIGHT Dir = iota
	DIR_LEFT
	DIR_UP
	DIR_DOWN
)

var AllDirs = []Dir{DIR_RIGHT, DIR_LEFT, DIR_UP, DIR_DOWN}

func (d Dir) String() string {
	switch d {
	case DIR_RIGHT:
		return "R"
	case DIR_LEFT:
		return "L"
	case DIR_UP:
		return "U"
	case DIR_DOWN:
		return "D"
	default:
		panic(d) // unreachable
	}
}

func (d Dir) Opposite() Dir {
	switch d {
	case DIR_RIGHT:
		return DIR_LEFT
	case DIR_LEFT:
		return DIR_RIGHT
	case DIR_UP:
		return DIR_DOWN
	case DIR_DOWN:
		return DIR_UP
	default:
		panic(d) // unreachable
	}
}

func (d Dir) UpdateTax(tax float64) float64 {
	switch d {
	case DIR_RIGHT:
		return tax + 2
	case DIR_LEFT:
		return tax - 2
	case DIR_UP:
		return tax / 2
	case DIR_DOWN:
		return tax * 2
	}
	panic(d) // unreachable
}
