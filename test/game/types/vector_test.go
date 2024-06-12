package types_test

import (
	"testing"

	"github.com/bbfh-dev/proglib-task_1-2/game/types"
)

func TestVector(t *testing.T) {
	vec := types.NewVector(0, 0)

	expected := types.NewVector(10, 0)
	vec = vec.Add(expected)
	if vec.X != expected.X || vec.Y != expected.Y {
		t.Fatalf("Failed to add. Expected %v; Got: %v", expected, vec)
	}

	if !vec.Equals(expected) {
		t.Fatalf("vec.Equals(). Expected %v; Got: %v", expected, vec)
	}
}
