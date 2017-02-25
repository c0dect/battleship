package battleship

import (
	"fmt"
	"io"
)

// Game represents the current state of any game.
type Game struct {
	Player1 *Player
	Player2 *Player

	P1Hits int
	P2Hits int
}

// PrintResult prints the result using the
// current state of the Game.
func (g *Game) PrintResult(w io.Writer) {
	g.Player1.printBoard(w)
	fmt.Fprintln(w)
	g.Player2.printBoard(w)
	fmt.Fprintln(w)
	fmt.Fprintf(w, "P1:%d\nP2:%d\n", g.P1Hits, g.P2Hits)

	switch {
	case g.P1Hits == g.P2Hits:
		fmt.Fprintln(w, "It is a draw")
		return
	case g.P1Hits > g.P2Hits:
		fmt.Fprintln(w, "Player 1 wins")
		return
	case g.P1Hits < g.P2Hits:
		fmt.Fprintln(w, "Player 2 wins")
		return
	}
}

// SimulateGame runs the game by using the missile coordinates
// defined by both players.
func (g *Game) SimulateGame() {
	// fire missiles for Player 1.
	for _, v := range g.Player1.MissileTargets {
		g.P1Hits += g.Player2.missileHitStatus(v)
	}

	// fire missiles for Player 2.
	for _, v := range g.Player2.MissileTargets {
		g.P2Hits += g.Player1.missileHitStatus(v)
	}

	//g.PrintResult()
}
