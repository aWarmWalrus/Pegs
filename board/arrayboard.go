package board

import (
    "fmt"
)

var DEBUG = false

type ArrayBoard struct {
    G [][]int
}

// ----- ArrayBoard Methods -----------------------------------------------------------

func (a *ArrayBoard) Init() {
    fullWidth := Width * 2 + 1
    fullHeight := Height * 2 + 1
    a.G = make([][]int, fullWidth)
    for i := 0; i < fullWidth; i++ {
        a.G[i] = make([]int, fullHeight)
        for j := 0; j < fullHeight; j++ {
            a.G[i][j] = 0
        }
    }
}

func (a *ArrayBoard) InitFrom(orig ArrayBoard) {
    fullWidth := Width * 2 + 1
    fullHeight := Height * 2 + 1
    a.G = make([][]int, fullWidth)
    for i := 0; i < fullWidth; i++ {
        a.G[i] = make([]int, fullHeight)
        for j := 0; j < fullHeight; j++ {
            val := orig.G[i][j]
            a.G[i][j] = val
        }
    }
}

func (a ArrayBoard) EqualTo(orig ArrayBoard) bool {
    for i := 0; i < Width * 2 + 1; i ++ {
        for j := 0; j < Width * 2 + 1; j++ {
            if a.G[i][j] != orig.G[i][j] {
                return false
            }
        }
    }
    return true
}

func (a ArrayBoard) InQueue(queue []ArrayBoard) bool {
    for _, b := range(queue) {
        if a.EqualTo(b) {return true}
    }
    return false
}

func (a ArrayBoard) Debug() {
    if DEBUG {
        a.Print()
    }
}

func (a ArrayBoard) Print() {
    fullW := Width * 2 + 1
    fullH := Height * 2 + 1
    fmt.Printf("Printing the board\n");
    for j := fullH - 1; j >= 0; j-- {
        for i := 0; i < fullW; i++ {
            fmt.Printf("%v ", a.G[i][j])
        }
        fmt.Printf("\n")
        if j == Height + 1 {
            fmt.Printf("---------------\n")
        }
    }
}

func (a *ArrayBoard) PutPeg(c Coord) error {
    x := c.X + Width;
    y := c.Y + Height;
    if ((x >= Width * 2 + 1) || (y >= Height * 2 + 1) || (x < 0) || (y < 0) || (a.G[x][y] == 1)) {
        return fmt.Errorf("err: PutPeg(%v, %v) : invalid coordinates", c.X, c.Y)
    }
    a.G[x][y] = 1
    return nil
}

func (a *ArrayBoard) RemovePeg(c Coord) error {
    x := c.X + Width;
    y := c.Y + Height;
    if ((x >= Width * 2 + 1) || (y >= Height * 2 + 1) || (x < 0) || (y < 0) || (a.G[x][y] == 0)) {
        return fmt.Errorf("err: RemovePeg(%v, %v) : invalid coordinates", c.X, c.Y)
    }
    a.G[x][y] = 0
    return nil
}

func (a *ArrayBoard) ExpandPeg(c Coord, direction string) error {
    if (a.G[c.X + Width][c.Y + Height] == 0) {
        return fmt.Errorf("err: ExpandPeg(%v, %v) - no peg at requested location", c, direction)
    }

    switch direction {
    case "DOWN":
        if err := a.RemovePeg(c); err != nil {
            return err
        } else if err := a.PutPeg(Coord{c.X, c.Y-1}); err != nil {
            return err
        } else if err := a.PutPeg(Coord{c.X, c.Y-2}); err != nil {
            return err
        }
    case "LEFT":
        if err := a.RemovePeg(c); err != nil {
            return err
        } else if err := a.PutPeg(Coord{c.X-1, c.Y}); err != nil {
            return err
        } else if err := a.PutPeg(Coord{c.X-2, c.Y}); err != nil {
            return err
        }
    case "RIGHT":
        if err := a.RemovePeg(c); err != nil {
            return err
        } else if err := a.PutPeg(Coord{c.X+1, c.Y}); err != nil {
            return err
        } else if err := a.PutPeg(Coord{c.X+2, c.Y}); err != nil {
            return err
        }
    default:
        return fmt.Errorf("err: ExpandPeg(%v, %v) - invalid direction", c, direction)
    }
    return nil
}

// Returns a slice of pegs in their pure coordinates
func (a ArrayBoard) WinMa() (bool, []Coord) {
    existingPegs := make([]Coord, 5)
    winning := true
    for i := 0; i < Width * 2 + 1; i++ {
        for j := 0; j < Height * 2 + 1; j++ {
            if a.G[i][j] == 1 {
                if j > Height {
                    winning = false
                }
                existingPegs = append(existingPegs, Coord{i - Width, j - Height})
            }
        }
    }
    return winning, existingPegs
}
