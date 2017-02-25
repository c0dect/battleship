package battleship_test

import (
	"reflect"
	"testing"

	"github.com/c0dect/battleship"
)

func TestCreatePlayer(t *testing.T) {
	targets := []battleship.Coordinates{
		battleship.Coordinates{0, 0},
		battleship.Coordinates{2, 2},
		battleship.Coordinates{1, 1},
	}
	player := &battleship.Player{
		Name: "TestPlayer",
		Board: [][]string{
			[]string{"B", "_", "_"},
			[]string{"_", "B", "_"},
			[]string{"_", "_", "B"},
		},
		MissileTargets: targets,
	}

	coordinates := []battleship.Coordinates{
		battleship.Coordinates{0, 0},
		battleship.Coordinates{2, 2},
		battleship.Coordinates{1, 1},
	}

	createdPlayer := battleship.CreatePlayer("TestPlayer", 3, coordinates, targets)

	if !reflect.DeepEqual(player, createdPlayer) {
		t.Fatalf("Player creation failure. Exp: %+v, Got: %+v", player, createdPlayer)
	}
}
