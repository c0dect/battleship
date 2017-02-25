package battleship

import (
	"fmt"
	"io"
	"strings"
)

// Player represents a single player of the game.
type Player struct {
	// Name represents the name of the player.
	Name string

	// Board represents the present battleship board of
	// the player at any given stage.
	//
	// _ is the initial state of a cell.
	// B is an alive battleship.
	// X is a missile hit.
	// O is a missile miss.
	Board [][]string

	// MissileTargets represents the target coordinates
	// to be fired on the opponent.
	MissileTargets []Coordinates
}

// CreatePlayer initializes the player object.
func CreatePlayer(name string, boardSize int, shipCoords, targetCoords []Coordinates) *Player {
	board := make([][]string, boardSize)

	// initialized an empty board.
	for i := 0; i < boardSize; i++ {
		board[i] = make([]string, boardSize)
		for j := 0; j < boardSize; j++ {
			board[i][j] = "_"
		}
	}

	// place the battleships.
	for _, v := range shipCoords {
		board[v.X][v.Y] = "B"
	}

	return &Player{
		Name:           name,
		Board:          board,
		MissileTargets: targetCoords,
	}
}

// PrintBoard prints the current board status
// for the given player.
func (p *Player) printBoard(w io.Writer) {
	fmt.Fprintln(w, p.Name)
	for i := 0; i < len(p.Board[0]); i++ {
		var row string
		for j := 0; j < len(p.Board[0]); j++ {
			row += p.Board[i][j] + " "
		}
		fmt.Fprintf(w, "%s\n", strings.TrimSuffix(row, " "))
	}
}

// MissileHitStatus for a player returns 1 if the player's
// battleship was hit by the missile passed in the coordinate.
func (p *Player) missileHitStatus(c Coordinates) int {
	if p.Board[c.X][c.Y] == "B" {
		p.Board[c.X][c.Y] = "X"
		return 1
	}
	p.Board[c.X][c.Y] = "O"
	return 0
}
