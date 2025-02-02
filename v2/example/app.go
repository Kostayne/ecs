package example

import (
	"fmt"

	"github.com/kostayne/ecs/v2/core"
)

func main() {
	ecs := core.MakeECS()

	// creating a movement system, it will change position components
	moveSys := MakeMovementSystem()

	// creating components
	posComp := MakePositionComponent(0, 0)
	velComp := MakeVelocityComponent(1, 1)

	// adding a new entity with provided components
	player := ecs.EntityStore.New(posComp, velComp)

	// adding movement system to ecs
	ecs.SystemStore.Add(moveSys)

	// getting a component to display
	plPos := (*player.GetOne("position")).(*PositionComponent)

	ecs.Setup()

	// main loop
	for i := 0; i < 5; i++ {
		// calling systems logic
		ecs.Process()

		// displaying results
		fmt.Printf("XY: (%v, %v)\n", plPos.X, plPos.Y)
	}

	ecs.Cleanup()
}
