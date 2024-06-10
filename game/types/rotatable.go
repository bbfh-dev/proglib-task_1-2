package types

type Rotatable interface {
	Rotation() Rotation
	SetRotation(Rotation) Rotatable
}

type Rotation float64

type Direction int

const (
	CLOCKWISE        Direction = 1
	COUNTERCLOCKWISE Direction = -1
)

func Rotate[T Rotatable](rotatable T, angle Rotation) T {
	rotatable = rotatable.SetRotation(rotatable.Rotation() + angle).(T)
	return rotatable
}
