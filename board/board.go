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

// ROW 5 CONDITIONS
var WinCond5 = []Coord{
	{-4, 0}, {-3, 0}, {-2, 0}, {-1, 0}, {0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0},
	{-1, -1}, {0, -1}, {1, -1}, {2, -1}, {3, -1},
	{-1, -2}, {0, -2}, {1, -2}, {2, -2},
	{0, -3}, {2, -3},
}

var Target5 = []Coord{
	{-2, 0}, {-1, 0}, {0, 0}, {1, 0}, {2, 0},
	{0, -1},
	{0, -2},
	{0, -3},
}

// ROW 4 CONDITIONS
var WinCond4 = []Coord{
	{-2, 0},
	{-2, -1},
	{-1, 0},
	{-1, -1},
	{0, 0},
	{0, -1},
	{1, 0},
	{2, 0},
}

// TARGET is the set of coordinates the pegs start at
var Target4 = []Coord{
	{-2, 0},
	{-1, 0},
	{0, -1},
	{1, 0},
	{2, 0},
}

// ROW 3 CONDITIONS
var WinCond3 = []Coord{
	{0, 0},
	{1, 0},
	{2, 0},
	{0, -1},
}

var Target3 = []Coord{
	{0, 0},
	{0, -1},
}
