package game_test

import (
	"testing"

	"github.com/bbfh-dev/proglib-task_1-2/game"
	"github.com/bbfh-dev/proglib-task_1-2/game/types"
)

func TestMovement(t *testing.T) {
	ship := game.NewSpaceShip(types.NewVector(12, 5), 0)
	ship.SetVelocity(types.NewVector(-7, 3))
	ship = types.Move(ship)

	expected := types.NewVector(5, 8)
	if !ship.Position().Equals(expected) {
		t.Fatalf("New ship position (%v) does not equal %v", ship.Position(), expected)
	}
}

func TestRotation(t *testing.T) {
	ship := game.NewSpaceShip(types.NewVector(0, 0), 0)

	ship = types.Rotate(ship, 45)
	expected := types.Rotation(45)
	if ship.Rotation() != expected {
		t.Fatalf("New ship rotation (%v) does not equal %v", ship.Rotation(), expected)
	}

	ship = types.Rotate(ship, -80)
	expected = types.Rotation(-35)
	if ship.Rotation() != expected {
		t.Fatalf("New ship rotation (%v) does not equal %v", ship.Rotation(), expected)
	}
}
