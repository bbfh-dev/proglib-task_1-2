package types

import "fmt"

type Vector struct {
	X float64
	Y float64
}

func NewVector(x float64, y float64) Vector {
	return Vector{X: x, Y: y}
}

func (vec Vector) String() string {
	return fmt.Sprintf("Vec(%v, %v)", vec.X, vec.Y)
}

func (vec Vector) Add(vec2 Vector) Vector {
	vec.X += vec2.X
	vec.Y += vec2.Y
	return vec
}

func (vec Vector) Equals(vec2 Vector) bool {
	return vec.X == vec2.X && vec.Y == vec2.Y
}
