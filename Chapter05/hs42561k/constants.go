package hs42561k

// Each bit translates to a pin, which is driven high or low
type Character byte

func (char Character) String() string {
	switch char {
	case Zero:
		return "0"
	case One:
		return "1"
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Dash:
		return "-"
	case E:
		return "E"
	case H:
		return "H"
	case L:
		return "L"
	case P:
		return "P"
	case Blank:
		return ""
	case Dot:
		return "."
	}

	return ""
}

const (
	Zero  Character = 0
	One   Character = 1
	Two   Character = 2
	Three Character = 3
	Four  Character = 4
	Five  Character = 5
	Six   Character = 6
	Seven Character = 7
	Eight Character = 8
	Nine  Character = 9
	Dash  Character = 10
	E     Character = 11
	H     Character = 12
	L     Character = 13
	P     Character = 14
	Blank Character = 15
	Dot   Character = 128
)
