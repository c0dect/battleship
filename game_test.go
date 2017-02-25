package battleship_test

import (
	"os"

	"github.com/c0dect/battleship"
)

func ExampleGame_SimulateGame() {
	p1 := &battleship.Player{
		Name: "Player1",
		Board: [][]string{
			[]string{"_", "_", "_", "_", "_"},
			[]string{"_", "B", "_", "_", "_"},
			[]string{"B", "_", "_", "B", "_"},
			[]string{"_", "_", "_", "_", "B"},
			[]string{"_", "_", "_", "B", "_"},
		},
		MissileTargets: []battleship.Coordinates{
			battleship.Coordinates{0, 1},
			battleship.Coordinates{4, 3},
			battleship.Coordinates{2, 3},
			battleship.Coordinates{3, 1},
			battleship.Coordinates{4, 1},
		},
	}
	p2 := &battleship.Player{
		Name: "Player2",
		Board: [][]string{
			[]string{"_", "B", "_", "_", "_"},
			[]string{"_", "_", "_", "_", "_"},
			[]string{"_", "_", "_", "B", "_"},
			[]string{"B", "_", "_", "_", "B"},
			[]string{"_", "B", "_", "_", "_"},
		},
		MissileTargets: []battleship.Coordinates{
			battleship.Coordinates{0, 1},
			battleship.Coordinates{0, 0},
			battleship.Coordinates{1, 1},
			battleship.Coordinates{2, 3},
			battleship.Coordinates{4, 3},
		},
	}

	game := &battleship.Game{
		Player1: p1,
		Player2: p2,
	}

	game.SimulateGame()
	game.PrintResult(os.Stdout)
	// Output:
	// Player1
	// O O _ _ _
	// _ X _ _ _
	// B _ _ X _
	// _ _ _ _ B
	// _ _ _ X _
	//
	// Player2
	// _ X _ _ _
	// _ _ _ _ _
	// _ _ _ X _
	// B O _ _ B
	// _ X _ O _
	//
	// P1:3
	// P2:3
	// It is a draw
}

//func TestGameScenarios(t *testing.T) {
//	var scenarios := []struct{
//
//	}
//}
