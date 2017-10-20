package main

import (
	b "Pegs/board"
	"flag"
	"fmt"
	"time"
)

var MOVES = []string{"DOWN", "RIGHT", "LEFT"}

const USING_ARRAYBOARD = false
const DEBUG = true

var deduping = flag.Bool("deduping", true, "True by default. If set to false, the program will not clean the queue for duplicate stages")
var row = flag.Int("row", 0, "0 by default. Sets the difficulty, between 3 and 5.")
var board = flag.String("board", "MAP", "Options: MAP, ARRAY, STRING (tbi)")
var maxdepth = flag.Int("maxdepth", 20, "The Max Depth")
var width = flag.Int("width", 4, "Board Width")
var height = flag.Int("height", 4, "Board Height")

//TODO
func debugPrint() {
}

//--------------------------------------------------------------------------------------

func handle(posQ map[uint64]bool, depth int) (int, error) {
	fmt.Printf("\n NEXT LEVEL OF DEPTH %v ---- Queue size: %v ----\n", depth, len(posQ))
	iTime := time.Now()
	if depth > (*maxdepth) {
		return 0, fmt.Errorf("handle: No solution found!")
	}
	win := false
	nextQueue := make(map[uint64]bool, len(posQ))
	symQ := make(map[uint64]bool, len(posQ)) // 1
	for bNum, _ := range posQ {
        board := b.IntBoard{bNum}
		winRes, nextPegs := board.WinMa()
		if winRes {
			fmt.Printf("Solution found:\n")
			board.Print()
			win = true
		}
		for _, tPeg := range nextPegs {
			for _, direction := range MOVES {
				dboard := b.IntBoard{}
				//fmt.Printf("board:\n %v\n", board)
				dboard.InitFrom(board) // copy from original board
				if err := dboard.ExpandPeg(tPeg, direction); err != nil {
					// don't do anything. maybe debugLog it
                // Only add the board if it's not in the symmetry
                // map already
                } else if _, ok := symQ[dboard.G]; !ok { // 2
					nextQueue[dboard.G] = true
                    symQ[dboard.Symmetry()] = true // 3
				}
			}
		}
	}
	if win {
		return depth, nil
	}
    fmt.Printf("Win: %v\n  -- time elapsed for processing: %v\n", win, time.Since(iTime).String())

    return handle(nextQueue, depth+1)
}

func solve(row int) (int, error) {
	fmt.Printf("Solving!\n")
	gameBoard := b.IntBoard{}
	if err := gameBoard.Init(row, *width, *height); err != nil {
		return 0, err
	}
	if USING_ARRAYBOARD {
		if err := gameBoard.PutPeg(b.Coord{0, row}); err != nil {
			return 0, err
		}
		if err := gameBoard.ExpandPeg(b.Coord{0, row}, "DOWN"); err != nil {
			return 0, err
		}
	}
	gameBoard.Print()
	positionMap := make(map[uint64]bool, 1)
	positionMap[gameBoard.G] = true
    
	res, err := handle(positionMap, 0)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func main() {
	// Initialize the flags
	flag.Parse()
	initTime := time.Now()
	res, err := solve(*row)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Min number of additional pegs for row %v: %v\n", *row, res+2)
	fmt.Printf(" -- Run duration: %v\n", time.Since(initTime).String())
}
