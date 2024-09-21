
package main

import "math/rand"

// play plays the game for the current player.
// + registers the player's move in the board.
// if the move is valid:
// + increases the turn.
func play() {
	// [" ", " ", " ", " ", " ", " ", " ", " ", " " ] -> cells slice
	//   0    1    2    3    4    5    6    7    8    -> indexes

	// /---+---+---\
	// | 0 | 1 | 2 |
	// +---+---+---+
	// | 3 | 4 | 5 |
	// +---+---+---+
	// | 6 | 7 | 8 |
	// \---+---+---/

	// pick a random move (very intelligent AI!)
	// it can play to the same position!
	lastPos = rand.Intn(maxTurns)

	// is it a valid move?
	if cells[lastPos] != emptyCell {
		wrongMove = true

		// skip the rest of the function from running
		return
	}

	// register the move: put the player's sign on the board
	cells[lastPos] = player

	// increment the current turns
	turn++
}

// switchPlayer switches to the next player
func switchPlayer() {
	// switch the player
	if player == player1 {
		player = player2
	} else {
		player = player1
	}
}
