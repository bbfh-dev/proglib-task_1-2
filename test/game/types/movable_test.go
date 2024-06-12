package types_test

import (
	"testing"

	"github.com/bbfh-dev/proglib-task_1-2/game/types"
)

type ExampleMovable struct {
	position types.Vector
	velocity types.Vector
}

func (test ExampleMovable) Position() types.Vector {
	return test.position
}

func (test ExampleMovable) Velocity() types.Vector {
	return test.velocity
}

func (test ExampleMovable) SetPosition(vec types.Vector) types.Movable {
	test.position = vec
	return test
}

func TestMovable(t *testing.T) {
	movable := ExampleMovable{
		position: types.NewVector(0, 0),
		velocity: types.NewVector(1, 0),
	}

	movable = types.Move(movable)
	expected := types.NewVector(1, 0)
	if !movable.Position().Equals(expected) {
		t.Fatalf("New movable position %v does not equal %v", movable.Position(), expected)
	}
}
