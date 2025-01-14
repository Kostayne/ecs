package example

import (
	"github.com/kostayne/ecs/core"
)

type MovementSystem struct{}

func (s *MovementSystem) GetType() string {
	return "sys_movement"
}

// Optional
func (s *MovementSystem) Setup(es *core.EntityStore)   {}
func (s *MovementSystem) Cleanup(es *core.EntityStore) {}

// Main logic
func (s *MovementSystem) Process(es *core.EntityStore) {
	finder := core.MakeFinder(es)
	entities := finder.Has("position").GetMany()

	for _, e := range entities {
		comp := *e.Get("position")
		pos := comp.(*PositionComponent)

		// one line version
		// pos := (*e.Get("position")).(*PositionComponent)

		pos.X += 1
		pos.Y += 2
	}

	// Or just iterate over all entities
	// for _, e := range es.Entities {
	// 	if e.Has("position", "velocity") {
	// 		comps := e.GetList("position", "velocity")

	// 		pos := (*comps[0]).(*PositionComponent)
	// 		vel := (*comps[1]).(*VelocityComponent)

	// 		pos.X += vel.X
	// 		pos.Y += vel.Y
	// 	}
	// }
}

func MakeMovementSystem() *MovementSystem {
	return &MovementSystem{}
}
