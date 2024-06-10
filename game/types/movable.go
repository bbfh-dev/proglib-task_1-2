package types

type Movable interface {
	Position() Vector
	Velocity() Vector
	SetPosition(Vector) Movable
}

func Move[T Movable](movable T) T {
	movable = movable.SetPosition(movable.Position().Add(movable.Velocity())).(T)
	return movable
}
