package main

import (
	"fmt"
	"math/rand"

	//	lib "n-puzzle/golib"
	gui "Gomoku/gui"
)

type coordinate struct {
	x uint8
	y uint8
}

type position struct { // alternatively uint8 (0 = unocupied), but memory waste
	occupied bool
	player   bool
}

type game struct {
	goban  [19][19]position
	player bool // whose move is it? (player 0 - black first)
	//	capture0	uint8				// capture 10 and win
	//	capture1	uint8				// capture 10 and win
	//	move		uint32				// how many moves have been played in total (is this desirable/necessary?)
	//	pass		bool				// was the last move a pass (if next move pass -> game over)
	//	last0		coordinate			// last move to check ko rule () for player 0
	//	last1		coordinate			// last move to check ko rule () for player 1

}

// type ai struct { 	/// merge with game struct?
//	aiplayer	bool	// is player 1 human or AI
//	hotseat		bool	// AI player only suggests moves, human must choose move
//	prescience	uint8	// how many moves in advance do we examine
// }

func InitializeGame() *game {
	g := game{}
	return &g
}

func RandomCoordinate() coordinate {
	x := uint8(rand.Intn(19))
	y := uint8(rand.Intn(19))
	random := coordinate{x, y}
	return random
}

func PlaceStone(coordinate coordinate, player bool, goban *[19][19]position) {
	goban[coordinate.x][coordinate.y].occupied = true
	goban[coordinate.x][coordinate.y].player = player
}

func RemoveStone(coordinate coordinate, goban *[19][19]position) {
	goban[coordinate.x][coordinate.y].occupied = false
}

// func GameLoop(g) {
// if AI:
//		suggest move
// if human or hotseat:
//		listen for mouse click 						//###### do first
//			find position/pass/new/exit clicked 		//###### do first
//			if pass, double pass end?
//			if reset, reset game with new config
// 		check if position is valid (if human, assume ai has aleady checked)
//			occupied?
//			rules
//				ko
//				double-three
// PlaceStone(coordinate, true, &g.goban)
// check if capture
//		remove captured
//		update game.captured struct
// check win
//		5 in a row
// 		all win conditions?
// update game. struct
//		player change
//		moves ++
// re-render ebiten with updated goban and stats
//	return err
// }

func main() {
	fmt.Println("Hello world!") ////////
	g := InitializeGame()
	// launch ebiten. render goban and game stats 		// # do first
	// GameLoop(g)

	PlaceStone(RandomCoordinate(), true, &g.goban)  ////////
	PlaceStone(RandomCoordinate(), true, &g.goban)  //////
	PlaceStone(RandomCoordinate(), true, &g.goban)  ///////
	PlaceStone(RandomCoordinate(), false, &g.goban) ////////
	PlaceStone(RandomCoordinate(), true, &g.goban)  ////////

	zero := coordinate{0, 0}    /////////
	RemoveStone(zero, &g.goban) /////////

	gui.RunEbiten()

	fmt.Println(g.goban)                // whole goban //////////
	fmt.Println(g.goban[0][0])          // one position ///////////
	fmt.Println(g.goban[0][0].occupied) // one position occupied /////////
	fmt.Println(g.goban[0][0].player)   // one position player  ///////////
	fmt.Println("Goodbye world!")       ////////
}

// Rules:
// a stone, or solidly connected group of stones of one color, is captured if all liberties are occupied
// only allowed to place a stone in a position with no liberties if it immediately captures
