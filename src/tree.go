package play

import (
	"fmt"
	"math"
	// play "Gomoku.play"
)

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

type node struct {
	id		int
	value   int
	goban   [19][19]position 
	player	bool
	// game     *play.Game
	children []*node
	bestMove *node
}

func newNode(id int, value int, newGoban *[19][19]position, newPlayer bool) *node {
	return &node{
		id:   	id,
		value:	value,	// change this to initialize to zero
		goban:	newGoban
		player: newPlayer
	}
	// node.board = &play.Game
	// return node
}

// Recursively finds node by ID, and then appends child to node.chilren
func addChild(node *node, parentID int, child *node) int {
	if node.id == parentID {
		node.children = append(node.children, child)
		return 1
	} else {
		for idx, _ := range node.children {
			current := node.children[idx]
			addChild(current, parentID, child)
		}
	}
	return 0
}

// Recursively generates every move for a board (to depth 3), assigns value, and adds to tree
func generateBoardsDepth(depth int8, current *node, goban *[19][19]position, player bool) {
	if depth > 3 {
		return
	}
	for y = 0; y < 19; y++ {
		for x = 0; x < 19; x++ {
			coordinate := {y, x}
			if isMoveValid2(coordinate, goban) == true {		// duplicate of isMoveValid w/o *game
				id += 1
				newGoban := placeStone(coordinate, player, goban)
				value := valueBoard(newGoban, player)
				child = newNode(id, value, newGoban, player)
				addChild(current, current.id, child) //
				generateBoardsDepth(depth+1, child, !player)
			}
			continue
		}
	}
}

func createTree(goban *[19][19]position, player bool) *node {
	var y		int8
	var x		int8
	var tmp		*node

	id := 1
	root := newNode(id, 10, g.goban, g.player)
	generateBoardsDepth(3, root, g.goban, g.player)
	return root
}

func printTree(parent *node) {
	current := parent
	fmt.Printf("\nparent: %d\n", current.id)
	for i := range current.children {
		child := current.children[i]
		fmt.Printf("child: %d", child.id)
		// put in a mutex/lock to wait until this range is done, and then call printTree for the child
	}
	for i := range current.children {
		current := current.children[i]
		printTree(current)
	}
}

func minimaxTree() {
	root := createTree()
	// printTree(root)
	// fmt.Println("-----")
	// alpha := float64(math.Inf(-1))
	// beta := float64(math.Inf(1))
	alpha := minInt
	beta := maxInt
	minimaxRecursive(root, 3, alpha, beta, false)	// for some reason, maximizingplayer has to be set to 'false' for this to work
	current := root
	for current.bestMove != nil {
		fmt.Println(current.id)
		current = current.bestMove
	}
	fmt.Println(current.id)
}


//  creates a tree, whose root is the goban
//  creates all possible moves/boards to depth _, calculates values, add to tree
//  applies minimax to tree, finds best move



	// addChild(root, 1, &node{id: 2, Value: 20})
	// addChild(root, 1, &node{id: 3, Value: 30})
	// addChild(root, 1, &node{id: 4, Value: 40})
	// addChild(root, 2, &node{id: 5, Value: 50})
	// addChild(root, 2, &node{id: 6, Value: 60})
	// addChild(root, 2, &node{id: 7, Value: 70})
	// addChild(root, 3, &node{id: 8, Value: 80})
	// addChild(root, 3, &node{id: 9, Value: 90})
	// addChild(root, 3, &node{id: 10, Value: 100})
	// addChild(root, 4, &node{id: 11, Value: 110})