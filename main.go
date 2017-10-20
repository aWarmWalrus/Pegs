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

func handle(posQ []b.MapBoard, depth int) (int, error) {
	fmt.Printf("\n NEXT LEVEL OF DEPTH %v ---- Queue size: %v ----\n", depth, len(posQ))
	iTime := time.Now()
	if depth > (*maxdepth) {
		return 0, fmt.Errorf("handle: No solution found!")
	}
	win := false
	nextQueue := []b.MapBoard{}
	for _, board := range posQ {
		winRes, nextPegs := board.WinMa()
		if winRes {
			fmt.Printf("Solution found:\n")
			board.Print()
			win = true
		}
		for _, tPeg := range nextPegs {
			for _, direction := range MOVES {
				dboard := b.MapBoard{}
				//fmt.Printf("board:\n %v\n", board)
				dboard.InitFrom(board) // copy from original board
				if err := dboard.ExpandPeg(tPeg, direction); err != nil {
					// don't do anything. maybe debugLog it
				} else { //if !dboard.InQueue(nextQueue) {
					nextQueue = append(nextQueue, dboard)
				}
			}
		}

	}
	if win {
		return depth, nil
	}
	if !(*deduping) {
		fmt.Printf("Win: %v\n  -- time elapsed for processing: %v\n", win, time.Since(iTime).String())

		return handle(nextQueue, depth+1)
	}
	fmt.Printf("Win: %v\nQueue Length before cleansing: %v\n", win, len(nextQueue))
	fmt.Printf("  -- time elapsed for processing: %v\n", time.Since(iTime).String())
	jTime := time.Now()
	cleanQueue := make([]b.MapBoard, 0, 10)
	for _, nb := range nextQueue {
		if !nb.InQueue(cleanQueue) {
			cleanQueue = append(cleanQueue, nb)
		}
	}
	fmt.Printf("Queue Length after cleansing: %v\n", len(cleanQueue))
	fmt.Printf(" -- time elapsed for cleansing: %v\n", time.Since(jTime).String())
	return handle(cleanQueue, depth+1)
}

func solve(row int) (int, error) {
	fmt.Printf("Solving!\n")
	gameBoard := b.MapBoard{}
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
	positionQueue := make([]b.MapBoard, 1)
	positionQueue[0] = gameBoard

	res, err := handle(positionQueue, 0)
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
