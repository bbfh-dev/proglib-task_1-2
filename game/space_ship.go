package game

import (
	"github.com/bbfh-dev/proglib-task_1-2/game/types"
)

type SpaceShip struct {
	position types.Vector
	velocity types.Vector
	rotation types.Rotation
}

func NewSpaceShip(position types.Vector, rotation types.Rotation) SpaceShip {
	return SpaceShip{
		position: position,
		velocity: types.NewVector(0, 0),
		rotation: rotation,
	}
}

func (ship *SpaceShip) SetVelocity(vec types.Vector) {
	ship.velocity = vec
}

// region Implement types.Movable interface

func (ship SpaceShip) Position() types.Vector {
	return ship.position
}

func (ship SpaceShip) Velocity() types.Vector {
	return ship.velocity
}

func (ship SpaceShip) SetPosition(vector types.Vector) types.Movable {
	ship.position = vector
	return ship
}

// region Implement types.Rotatable interface

func (ship SpaceShip) Rotation() types.Rotation {
	return ship.rotation
}

func (ship SpaceShip) SetRotation(rotation types.Rotation) types.Rotatable {
	ship.rotation = rotation
	return ship
}

// endregion
