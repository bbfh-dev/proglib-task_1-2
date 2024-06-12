package types_test

import (
	"testing"

	"github.com/bbfh-dev/proglib-task_1-2/game/types"
)

type ExampleRotatable struct {
	rotation types.Rotation
}

func (test ExampleRotatable) Rotation() types.Rotation {
	return test.rotation
}

func (test ExampleRotatable) SetRotation(rotation types.Rotation) types.Rotatable {
	test.rotation = rotation
	return test
}

func TestRotation(t *testing.T) {
	rotatable := ExampleRotatable{
		rotation: 0,
	}

	rotatable = types.Rotate(rotatable, 45)
	expected := types.Rotation(45)
	if rotatable.Rotation() != expected {
		t.Fatalf("New rotatable rotation %v does not equal %v", rotatable.Rotation(), expected)
	}

	rotatable = types.Rotate(rotatable, -30)
	expected = types.Rotation(45 - 30)
	if rotatable.Rotation() != expected {
		t.Fatalf("New rotatable rotation %v does not equal %v", rotatable.Rotation(), expected)
	}
}
