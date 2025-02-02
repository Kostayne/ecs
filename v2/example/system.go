package example

import (
	"time"

	"github.com/kostayne/ecs/v2/core"
)

// Use SystemBase to reduce boilerplate
type MovementSystem struct {
	core.SystemBase
}

// Main logic
func (s *MovementSystem) Process(es *core.EntityStore, dt time.Duration) {
	finder := core.MakeFinder(es)
	entities := finder.Has("position").GetMany()

	for _, e := range entities {
		comp := *e.GetOne("position")
		pos := comp.(*PositionComponent)

		// one line version
		// pos := (*e.Get("position")).(*PositionComponent)

		pos.X += 1
		pos.Y += 2
	}
}

func MakeMovementSystem() *MovementSystem {
	return &MovementSystem{
		// Define system params here
		SystemBase: *core.MakeSystemBase("sys_movement", 0, 0),
	}
}
