package board

import (
    "fmt"
)

type IntBoard struct {
    G int
    Sym int
}

func (i *IntBoard) Init(row int, width int, height int) error {
}

func (i *IntBoard) InitFrom(other IntBoard) {
}

func (l IntBoard) EqualTo(orig IntBoard) bool {
}

func (i IntBoard) InQueue(queue []IntBoard) bool {
}

func (i IntBoard) Debug() {
}

func (i IntBoard) Print() {
}

func (i *IntBoard) PutPeg(c Coord) error {
}

func (i *IntBoard) RemovePeg(c Coord) error {
}

// Optimization: LEFT expansions should never happen on the right
// side, and RIGHT expansion should never happen on the left
func (i *IntBoard) ExpandPeg(c Coord, direction string) error {
}
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

func (i IntBoard) WinMa() (bool, []Coord) {

}
