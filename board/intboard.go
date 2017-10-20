package board

import (
    "fmt"
)

type IntBoard struct {
    G uint64
}

// width x height must be less than 64
// which only leaves possibilities: 
//  - 10 x 6
//  - 9 x 7 <- let's start with this
//  - 8 x 8
var xmax = 4 // effectively 9 width
var ymax = 6 // effectively 7 height
var maxWidth = uint64(xmax * 2 + 1)

var winBitmap uint64
var targetInt uint64

func coordToBit(c Coord) uint64 {
    return uint64(-c.Y * (xmax * 2 + 1) + (c.X + xmax))
}

func bitToCoord(b uint64) Coord {
    x := int(b) % int(maxWidth) - xmax
    y := -(int(b) / int(maxWidth))
    return Coord{x, y}
}

func mapToInt(coords []Coord) uint64 {
    var g uint64 // map in int form
    for _, c := range coords {
         g = g | (1 << coordToBit(c))
    }
    return g
}

func (i IntBoard) Symmetry() uint64 {
    bit := uint64(0)
    f := uint64(0)
    for bit = 0; bit < 64; bit++ {
        rem := bit % maxWidth // remainder
        uBit := bit - rem + (maxWidth - rem) - 1
        s1 := i.G & (1 << bit)
        s2 := s1 >> bit
        s3 := s2 << uBit
        f = f | s3
    }
    return f
}

func (i *IntBoard) Init(row int, width int, height int) error {
    i.G = 0
    // we're gonna... ignore width and height for now haha)
	switch row {
	case 3:
		targetInt = mapToInt(Target3)
		winBitmap = mapToInt(WinCond3)
	case 4:
		targetInt = mapToInt(Target4)
		winBitmap = mapToInt(WinCond4)
	case 5:
		targetInt = mapToInt(Target5)
		winBitmap = mapToInt(WinCond5)
	default:
		return fmt.Errorf("Bad row param: %v", row)
	}

    i.G = targetInt
    return nil
}

func (i *IntBoard) InitFrom(other IntBoard) {
    i.G = other.G
}

func (i IntBoard) EqualTo(orig IntBoard) bool {
    return i.G == orig.G
}

func (i IntBoard) InMap(m map[uint64]bool) bool {
    _, ok := m[i.G]
    return ok
}

func (i IntBoard) Debug() {
    if DEBUG {
        i.Print()
    }
}

func (i IntBoard) Print() {
	fmt.Printf(" ---- MAP (%v x %v) ----\n", xmax*2+1, ymax+1)

    bit := uint64(0)
    for y := 0; y >= -ymax; y-- {
        fmt.Printf("   ")
        for x := -xmax; x <= xmax; x++ {
            if (i.G & (1 << bit)) != 0 {
                fmt.Printf("x ")
            } else {
                fmt.Printf("o ")
            }
            bit++
        }
        fmt.Printf("\n")
    }
    fmt.Printf(" ----------------------\n")
}

func (i *IntBoard) PutPeg(c Coord) error {
    b := coordToBit(c)
    if (i.G & (1 << b)) != 0 {
        return fmt.Errorf("err: PutPeg(%v): Peg already exists")
    }
	if (c.X > xmax) || (c.Y < -ymax) || (c.X < -xmax) || (c.Y > 0) {
		return fmt.Errorf("err: PutPeg(%v): Peg out of bounds")
	}
    i.G = i.G | (1 << b)
    return nil
}

func (i *IntBoard) RemovePeg(c Coord) error {
    b := coordToBit(c)
    if (i.G & (1 << b)) == 0 {
        return fmt.Errorf("err: PutPeg(%v): Peg doesn't exist")
    }
	if (c.X > xmax) || (c.Y < -ymax) || (c.X < -xmax) || (c.Y > 0) {
        fmt.Printf("Remove Peg Error??\n")
		return fmt.Errorf("err: PutPeg(%v): Peg out of bounds")
	}
    i.G = i.G &^ (1 << b)
    return nil
}

// Optimization: LEFT expansions should never happen on the right
// side, and RIGHT expansion should never happen on the left
func (i *IntBoard) ExpandPeg(c Coord, direction string) error {
	switch direction {
	case "DOWN":
		if err := i.RemovePeg(c); err != nil {
			return err
		} else if err := i.PutPeg(Coord{c.X, c.Y - 1}); err != nil {
			return err
		} else if err := i.PutPeg(Coord{c.X, c.Y - 2}); err != nil {
			return err
		}
	case "LEFT":
		if c.X > 0 {
			return fmt.Errorf("Left expansions on the right side are unnecessary")
		}
		if err := i.RemovePeg(c); err != nil {
			return err
		} else if err := i.PutPeg(Coord{c.X - 1, c.Y}); err != nil {
			return err
		} else if err := i.PutPeg(Coord{c.X - 2, c.Y}); err != nil {
			return err
		}
	case "RIGHT":
		if c.X < 0 {
			return fmt.Errorf("Right expansions on the left side are unnecessary")
		}
		if err := i.RemovePeg(c); err != nil {
			return err
		} else if err := i.PutPeg(Coord{c.X + 1, c.Y}); err != nil {
			return err
		} else if err := i.PutPeg(Coord{c.X + 2, c.Y}); err != nil {
			return err
		}
	default:
		return fmt.Errorf("err: ExpandPeg(%v, %v) - invalid direction", c, direction)
	}
	return nil
}

func (i IntBoard) WinMa() (bool, []Coord) {
    coords := make([]Coord, 0, 32)
    var m uint64
    for m = 0; m < 64; m++ {
        if (i.G & (1 << m)) != 0 {
            coords = append(coords, bitToCoord(m))
        }
    }

    return (i.G & winBitmap) == 0, coords
}
