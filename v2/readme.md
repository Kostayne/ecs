# Kostayne ECS v2
![GitHub Tag](https://img.shields.io/github/v/tag/kostayne/ecs?label=version)
![GitHub License](https://img.shields.io/github/license/kostayne/ecs)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/kostayne/ecs/go.yml)
![GitHub top language](https://img.shields.io/github/languages/top/kostayne/ecs?style=flat&logo=go&logoColor=white&logoSize=20px&label=Pure%20go)

This package provides a basic implementation of Entity Component System pattern. ECS separates logic from data, which makes the app more scalable and flexible. It's ideal for complex games or simulations.

### Definitions

**Entity** - a unique object in the world, it has no any logic.

**Component** - a piece of data that can be attached to an entity.

**System** - a set of logic that can be applied to the entity.

## TOC
<!-- - [Definitions](#definitions) -->
- [Usage](#usage)
    - [Components](#define-components)
    - [Systems](#define-systems)
    - [Main loop](#start-the-app)
- [Wiki](#wiki)
	- [ECS](#ecs)
	- [EntityStore](#entitystore)
		- [Manage entities](#manage-entities)
		- [Manage components](#manage-components)
	- [Finder](#finder)
		- [Constructor](#----finder-constructor----)
		- [Methods](#----finder-methods----)
			- [Has](#finderhascomponents-string-finder)
			- [Where](#finderwherepredicate-funcentity-bool-finder)
			- [GetOne](#findergetone-entity)
			- [GetMany](#findergetmany-entity)

## Usage
[Full example code is here!](https://github.com/kostayne/ecs/tree/main/example)

To use ECS, you need to define components and systems, then create a main loop.

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
```

The code above reduces boilerplate with the SystemBase struct, here is a [more explicit version](https://github.com/kostayne/ecs/tree/main/example/system_explicit.go) without it.

### Start the app
Finally, we'll start the app (app.go):

```go
package example

import (
	"fmt"

	"github.com/kostayne/ecs/v2/core"
)

func main() {
	ecs := core.MakeECS()

	// create component & system instances
	moveSys := MakeMovementSystem()
	posComp := MakePositionComponent(0, 0)

	// add them to ecs
	player := ecs.EntityStore.New(posComp)
	ecs.SystemStore.Add(moveSys)

	ecs.Setup()

	// run the main loop
	for i := 0; i < 5; i++ {
		ecs.Process()
		fmt.Printf("XY: (%v, %v)\n", posComp.X, posComp.Y)
	}

	ecs.Cleanup()
}
```

## Wiki
To find out more, see the [documentation](https://pkg.go.dev/github.com/kostayne/ecs/core).

### ECS

ECS is a core data structure that holds all entities and their components.

```go
type ECS struct {
	EntityStore EntityStore
	SystemStore SystemStore
}
```

### EntityStore

Use entity store to manage entities.


#### Manage entities
```go
entity := ecs.EntityStore.New()
ecs.EntityStore.Get(entity.Id())
ecs.EntityStore.Remove(entity.Id())
ecs.EntityStore.GetAll(entity.Id())
```

#### Manage components
```go
ecs.EntityStore.GetById(entity.Id())
ecs.EntityStore.AddTo(entity.Id(), MakePositionComponent(0, 0))
ecs.EntityStore.RemoveFrom(entity.Id(), "position")
```

### Finder

Finder is a helper that allows to find entities by components or arbitrary criteria.

### --- Finder Interface ---

```go
type FinderI interface {
	Get() Entity
	GetMany() []Entity
	Has(components ...string) FinderI
	Where(predicate func(Entity) bool) FinderI
}
```

#### --- Finder Constructor ---

```go
ecs := core.MakeECS()

finder := core.MakeFinder(&ecs.SystemStore)
```

#### --- Finder Methods ---

#### Finder.Has(components ...string) FinderI
Returns a finder with entities that have provided components.

```go
entities := finder.Has("position", "velocity").GetMany()
```

#### Finder.Where(predicate func(Entity) bool) FinderI
Returns a finder with entities that match provided predicate.

```go
func isEntityOnTheRight(e *Entity) bool {
	pos := (*e.GetOne("position")).(*PositionComponent)
	return pos.X > 0
}

entities := finder.Where(isEntityOnTheRight).GetMany()
```

#### Finder.GetOne() *Entity
Returns a single matched entity.

```go
player := finder.Has("character_controller").GetOne()
```

#### Finder.GetMany() []Entity
Returns a list of matched entities.
```
weapons := finder.Has("weapon").GetMany()
```