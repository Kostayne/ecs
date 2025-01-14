package example

import (
	"fmt"

	"github.com/kostayne/ecs/core"
)

func main() {
	ecs := core.MakeECS()

	// adding systems
	moveSys := MakeMovementSystem()
	ecs.SystemStore.Add(moveSys)

	// creating components
	posComp := MakePositionComponent(0, 0)
	velComp := MakeVelocityComponent(1, 1)

	// creating new entities with components
	player := ecs.EntityStore.AddNew(posComp, velComp)

	// setup
	ecs.Setup()
	plPos := (*player.Get("position")).(*PositionComponent)

	// main loop
	for i := 0; i < 5; i++ {
		ecs.Process()
		fmt.Printf("XY: (%v, %v)\n", plPos.X, plPos.Y)
	}

	ecs.Cleanup()
}
