# Kostayne ECS
This package provides a basic implementation of Entity Component System pattern. ECS separates logic from data, which makes the app more scalable and flexible. It's ideal for complex games.

### Definitions

**Entity** - a unique object in the world, it has no any logic.

**Component** - a piece of data that can be attached to an entity.

**System** - a set of logic that can be applied to the entity.

## Table of content
<!-- - [Definitions](#definitions) -->
- [Usage](#usage)
    - [Components](#define-components)
    - [Systems](#define-systems)
    - [Main loop](#start-the-app)

## Usage

### Define components

First, we'll define a simple component (component.go):

```go
package example

type PositionComponent struct {
	X float64
	Y float64
}

func (c *PositionComponent) Type() string {
	return "position"
}

func MakePositionComponent(x, y float64) *PositionComponent {
	return &PositionComponent{
		X: x,
		Y: y,
	}
}
```

### Define systems
Next, we'll define a simple system (system.go):

```go
package example

import (
	"github.com/kostayne/ecs/core"
)

type MovementSystem struct{}

func (s *MovementSystem) GetType() string {
	return "sys_movement"
}

// Optional
func (s *MovementSystem) Setup(es *core.EntityStore) {}
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
}

func MakeMovementSystem() *MovementSystem {
	return &MovementSystem{}
}
```

### Start the app
Finally, we'll start the app (app.go):

```go
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

	// creating new entities with components
	player := ecs.EntityStore.AddNew(posComp)

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
```

## Wiki

### Entity

Entity meant for working with components as a single object, it has unique id.

#### Structure
```go
type Entity struct {
	id         EntityID
	Components []Component `json:"components"`
}
```

#### Entity Methods

#### Entity.Id()
```
Returns entity id
```

#### Entity.Add(components ...Component)
```
Adds components to the entity
```

#### Entity.Remove(components ...Component)
```
Removes components from the entity
```

#### Entity.Has(componentTypes ...string) bool
```
Returns true if entity has provided components
```

#### Entity.Get(componentType string) *Component
```
Returns a component by provided type
```

#### Entity.GetList(componentTypes ...string) []*Component
```	
Returns a list of components by provided types
```
