package example

import (
	"time"

	"github.com/kostayne/ecs/v2/core"
)

// _EXP stands for explicit
type MovementSystem_EXP struct{}

func (s *MovementSystem_EXP) Type() string {
	return "sys_movement"
}

// System lifecycle
func (s *MovementSystem_EXP) Setup(es *core.EntityStore)   {}
func (s *MovementSystem_EXP) Cleanup(es *core.EntityStore) {}

func (s *MovementSystem_EXP) Priority() int   { return 0 }
func (s *MovementSystem_EXP) Frequency() uint { return 0 }

// Main logic
func (s *MovementSystem_EXP) Process(es *core.EntityStore, dt time.Duration) {
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

func MakeMovementSystemEXP() *MovementSystem {
	return &MovementSystem{}
}
