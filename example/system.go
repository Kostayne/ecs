package example

import (
	"time"

	"github.com/kostayne/ecs/core"
)

type MovementSystem struct{}

func (s *MovementSystem) GetType() string {
	return "sys_movement"
}

// System lifecycle
func (s *MovementSystem) Setup(es *core.EntityStore)   {}
func (s *MovementSystem) Cleanup(es *core.EntityStore) {}

func (s *MovementSystem) GetPriority() int   { return 0 }
func (s *MovementSystem) GetFrequency() uint { return 0 }

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
