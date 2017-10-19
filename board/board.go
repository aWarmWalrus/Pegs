package board

type Board interface {
	Init()
	InitFrom(a Board)
	EqualTo(orig Board) bool
	InQueue(queue []Board) bool
	Debug()
	Print()
	PutPeg(c Coord) error
	RemovePeg(c Coord) error
	ExpandPeg(c Coord, direction string) error
	WinMa() (bool, []Coord)
}

const (
	ARRAYBOARD = iota
	LISTBOARD
)

type Problem struct {
	Target     []Coord
	Constraint []Coord
}

type Coord struct {
	X int
	Y int
}

const Width = 3
const Height = 4
