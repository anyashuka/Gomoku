package play

//import (
//	lib "Gomoku/golib"
//	gui "Gomoku/GUI"
//)

type coordinate struct {
	y int8
	x int8
}

type position struct { // alternatively uint8 (0 = unocupied), but memory waste
	occupied bool
	player   bool
}

type align5 struct { //winning move for checking if opponent breaks it in the next move
	aligned5 bool
	winner   bool /// rm?
	winmove  coordinate
}

type game struct {
	goban    [19][19]position
	player   bool   // whose move is it? (player 0 - black first)
	capture0 uint8  // capture 10 and win
	capture1 uint8  // capture 10 and win
	align5   align5 // one player has aligned 5, however it can be broken. The other player must break it or lose.
	// move		uint32				// how many moves have been played in total (is this desirable/necessary?)
}

// type ai struct { 	/// merge with game struct?
//	aiplayer	bool	// is player 1 human or AI
//	hotseat		bool	// AI player only suggests moves, human must choose move
//	prescience	uint8	// how many moves in advance do we examine
// }

func InitializeGame() *game {
	g := game{}
	// g.player = true ///// rm, just to test
	return &g
}

func SwapPlayers(player bool) bool {
	if player == false {
		return true
	} else {
		return false
	}
}

func GameLoop(g *game) {
	validated := false
	coordinate := RandomCoordinate() /////
	for i := 0; i < 10000; i++ {     //moves
		validated, coordinate = PlaceRandomIfValid(g)
		if validated == true {
			Capture(coordinate, g)
			DumpGoban(&g.goban) //////
			CheckWin(coordinate, g)
			g.player = SwapPlayers(g.player)
		}
	}
	// update game.moves ++
	//	return err
}

func Play() {
	g := InitializeGame()
	GameLoop(g)

	// /// Test DoubleFree
	// zero := coordinate{0, 0}  /////////
	// PlaceIfValid(zero, g)     /////////
	// three := coordinate{1, 1} /////////
	// PlaceIfValid(three, g)    /////////
	// three = coordinate{3, 5}  /////////
	// PlaceIfValid(three, g)    /////////
	// three = coordinate{3, 4}  /////////
	// PlaceIfValid(three, g)    /////////
	// // g.player = true           //////
	// // three = coordinate{3, 2}  ///////// one of the three-aligned obstructed, therefore next move legal.
	// // PlaceIfValid(three, g)    /////////
	// // g.player = false          /////
	// g.player = true          //////
	// three = coordinate{3, 1} /////////
	// PlaceIfValid(three, g)   /////////
	// g.player = false         /////
	// three = coordinate{3, 3} ///////// Double Three, should be rejected
	// PlaceIfValid(three, g)   ///////// Double Three, should be rejected

	// // launch ebiten. render goban and game stats 		// # do first
	// // GameLoop(g)

	// /// Test Place stone
	// PlaceIfValid(RandomCoordinate(), g)
	// PlaceIfValid(RandomCoordinate(), g)
	// PlaceIfValid(RandomCoordinate(), g)
	// PlaceIfValid(RandomCoordinate(), g)
	// PlaceIfValid(RandomCoordinate(), g)
	// PlaceIfValid(RandomCoordinate(), g)
	// // PlaceStone(RandomCoordinate(), true, &g.goban)  ////////
	// // PlaceStone(RandomCoordinate(), true, &g.goban)  //////
	// // PlaceStone(RandomCoordinate(), true, &g.goban)  ///////
	// // PlaceStone(RandomCoordinate(), false, &g.goban) ////////
	// // PlaceStone(RandomCoordinate(), true, &g.goban)  ////////

	// RemoveStone(zero, &g.goban) /////////

	// coordinate, elapsed := ai(g)
	// PlaceIfValid(coordinate, g) //////
	// fmt.Println(elapsed)        // tme taken by ai //////////

	// gui.RunEbiten()
	// DumpGoban(&g.goban)/////
}