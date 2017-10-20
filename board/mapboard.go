package board

import (
	"fmt"
)

// y = 0 is the highest y value. all y values should be lower
// wincond is the set of coordinates we want to be clear of

var wincond, target []Coord
var width, height int

type MapBoard struct {
	G map[Coord]bool
}

// Depends on width and height, declared in board.go
// Max width of 16. Uses hexadecimal
// Each row in the grid is a hexadecimal number
/*
func handleBuilder(l MapBoard) int {
	functionalX := c.X - width
	functionalY := c.Y - height
	maxX := width * 2 + 1
	maxY := height + 1
	for _, := range l.G {

	}
	return
}
*/

func (l *MapBoard) Init(row int, width int, height int) error {
	l.G = make(map[Coord]bool)
	switch row {
	case 3:
		target = Target3
		wincond = WinCond3
	case 4:
		target = Target4
		wincond = WinCond4
	case 5:
		target = Target5
		wincond = WinCond5
	default:
		return fmt.Errorf("Bad row param: %v", row)
	}
	width = width
	height = height
	for _, v := range target {
		l.G[v] = true
	}
	return nil
}

func (l *MapBoard) InitFrom(other MapBoard) {
	l.G = make(map[Coord]bool)
	for k, v := range other.G {
		l.G[k] = v
	}
}

func (l MapBoard) EqualTo(orig MapBoard) bool {
	// this shouldn't really happen but can't be too sure
	if len(l.G) != len(orig.G) {
		return false
	}
	for k, _ := range orig.G {
		_, ok := l.G[k]
		if !ok {
			return false
		}
	}
	return true
}

func (l MapBoard) InQueue(queue []MapBoard) bool {
	for _, m := range queue {
		if l.EqualTo(m) {
			return true
		}
	}
	return false
}

func (m MapBoard) Debug() {
	if DEBUG {
		m.Print()
	}
}

func (m MapBoard) Print() {
	fmt.Printf(" ---- MAP (%v x %v) ----\n", width*2+1, height+1)

	for j := 0; j >= -height; j-- {
		fmt.Printf("   ")
		for i := -width; i <= width; i++ {
			if _, ok := m.G[Coord{i, j}]; ok {
				fmt.Printf("X ")
			} else {
				fmt.Printf("0 ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf(" ---------------------\n")
}

func (m *MapBoard) PutPeg(c Coord) error {
	if _, ok := m.G[c]; ok {
		return fmt.Errorf("err: PutPeg(%v): Peg already exists")
	}
	if (c.X > width) || (c.Y > height) || (c.X < -width) || (c.Y < -height) {
		return fmt.Errorf("err: PutPeg(%v): Peg out of bounds")
	}
	m.G[c] = true
	return nil
}

func (m *MapBoard) RemovePeg(c Coord) error {
	if _, ok := m.G[c]; !ok {
		return fmt.Errorf("err: PutPeg(%v): Peg doesn't exist")
	}
	delete(m.G, c)
	return nil
}

// Optimization: LEFT expansions should never happen on the right
// side, and RIGHT expansion should never happen on the left
func (m *MapBoard) ExpandPeg(c Coord, direction string) error {
	switch direction {
	case "DOWN":
		if err := m.RemovePeg(c); err != nil {
			return err
		} else if err := m.PutPeg(Coord{c.X, c.Y - 1}); err != nil {
			return err
		} else if err := m.PutPeg(Coord{c.X, c.Y - 2}); err != nil {
			return err
		}
	case "LEFT":
		if c.X > 0 {
			return fmt.Errorf("Left expansions on the right side are unnecessary")
		}
		if err := m.RemovePeg(c); err != nil {
			return err
		} else if err := m.PutPeg(Coord{c.X - 1, c.Y}); err != nil {
			return err
		} else if err := m.PutPeg(Coord{c.X - 2, c.Y}); err != nil {
			return err
		}
	case "RIGHT":
		if c.X < 0 {
			return fmt.Errorf("Right expansions on the left side are unnecessary")
		}
		if err := m.RemovePeg(c); err != nil {
			return err
		} else if err := m.PutPeg(Coord{c.X + 1, c.Y}); err != nil {
			return err
		} else if err := m.PutPeg(Coord{c.X + 2, c.Y}); err != nil {
			return err
		}
	default:
		return fmt.Errorf("err: ExpandPeg(%v, %v) - invalid direction", c, direction)
	}
	return nil
}

func (m MapBoard) WinMa() (bool, []Coord) {
	keys := make([]Coord, 0, len(m.G))
	for k, _ := range m.G {
		keys = append(keys, k)
	}

	for _, c := range wincond {
		if _, ok := m.G[c]; ok {
			return false, keys
		}
	}
	return true, keys
}
