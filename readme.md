# Kostayne ECS
![GitHub License](https://img.shields.io/github/license/kostayne/ecs)

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
	- [Finder](#finder)
		- [Constructor](#----finder-constructor----)
		- [Methods](#----finder-methods----)
			- [Has](#finderhascomponents-string-finder)
			- [Where](#finderwherepredicate-funcentity-bool-finder)
			- [GetOne](#findergetone-entity)
			- [GetMany](#findergetmany-entity)
	- [Entity](#entity)
		- [Structure](#entity-structure)
		- [Methods](#entity-methods)
			- [Get](#entitygetcomponenttype-string-component)
			- [GetList](#entitygetlistcomponenttypes-string-component)
			- [Has](#entityhascomponenttypes-string-bool)
			- [Add](#entityaddcomponents-component)
			- [Remove](#entityremovecomponents-component)
	- [System](#system)
		- [Interface](#system-interface)
		- [Methods](#system-methods)
			- [Setup](#systemsetupentitystore-entitystore)
			- [Process](#systemprocessentitystore-entitystore)
			- [Cleanup](#systemcleanupentitystore-entitystore)
	- [ECS Core](#ecs-core)
		- [Constructor](#----ecs-constructor----)
		- [Methods](#----ecs-core-methods----)
			- [Setup](#ecssetup)
			- [Process](#ecsprocess)
			- [Cleanup](#ecscleanup)
	- [EntityStore](#entitystore)
		- [Methods](#----entity-store-methods----)
			- [New](#entitystorenewcomponents-component-entity)
			- [Remove](#entitystoreremoveentity-entity)
			- [AddTo](#entitystoreaddtoentity-entity-components-component)
			- [RemoveFrom](#entitystoreremovefromentity-entity-componenttypes-string)
			- [GetComponentsById](#entitystoregetcomponentsbyidentityid-entityid-component)
	- [SystemStore](#system-store)
		- [Methods](#----system-store-methods----)
			- [Add](#systemstoreaddsystem-system)
			- [Remove](#systemstoreremovetypename-string)
			- [Get](#systemstoregettypename-string-system)
			- [GetAll](#systemstoregetall-system)
			- [GetPriority](#systemstoregetpriority-_systempriority)
			- [GetLastCallTime](#systemstoregetlastcalltime-mapstringtimetime)

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

// Hooks
func (s *MovementSystem) Setup(es *core.EntityStore) {}
func (s *MovementSystem) Cleanup(es *core.EntityStore) {}

// Params
func (s *MovementSystem) GetPriority() int   { return 0 }
func (s *MovementSystem) GetFrequency() uint { return 0 }

// Main logic
func (s *MovementSystem) Process(es *core.EntityStore, dt time.Duration) {
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

### Finder

Finder is a helper that allows to find entities by components or arbitrary criteria.

### --- Finder Interface ---

```go
type FinderI interface {
	Get() Entity
	GetMany() []Entity
	Has(components ...string) *FinderI
	Where(predicate func(*Entity) bool) *FinderI
}
```

#### --- Finder Constructor ---

```go
ecs := core.MakeECS()

finder := core.MakeFinder(&ecs.SystemStore)
```

#### --- Finder Methods ---

#### Finder.Has(components ...string) *Finder
Returns a finder with entities that have provided components

```go
entities := finder.Has("position", "velocity").GetMany()
```

#### Finder.Where(predicate func(*Entity) bool) *Finder
Returns a finder with entities that match provided predicate

```go
func isEntityOnTheRight(e *Entity) bool {
	pos := (*e.Get("position")).(*PositionComponent)
	return pos.X > 0
}

entities := finder.Where(isEntityOnTheRight).GetMany()
```

#### Finder.GetOne() *Entity
Returns a single matched entity

```go
player := finder.Has("character_controller").GetOne()
```

#### Finder.GetMany() []*Entity
Returns a list of matched entities
```
weapons := finder.Has("weapon").GetMany()
```

### Entity

Entity meant for working with components as a single object, it has unique id.

#### --- Structure ---
```go
type Entity struct {
	id         EntityID
	Components []Component `json:"components"`
}
```

#### --- Entity Methods ---

#### Entity.Id()
Returns entity id

```go
id := entity.Id()
```

#### Entity.GetOne(componentType string) *Component
Returns a component by provided type, may return nil

```go
pos := (*e.Get("position")).(*PositionComponent)
```

#### Entity.GetList(componentTypes ...string) []*Component
Returns a list of components by provided types

```go	
comps := entity.GetList("position", "velocity")
```

#### Entity.Has(componentTypes ...string) bool
Returns true if entity has provided components

```
hasPos := entity.Has("position")
```

#### Entity.Add(components ...Component)
Adds components to the entity

```go
entity.Add("position", "velocity")
```

#### Entity.Remove(components ...Component)
Removes components from the entity

```go
entity.Remove("position", "velocity")
```

### Component
Component stores specific data that describes an entity state like position, velocity, etc.

#### --- Structure ---
Component does not have any structure by default, it's up to you to define it.

```go
type Component struct {
	// anything you want!
}
```

#### --- Component Interface ---
The only thing you need to implement is the Component.Type() method. The return value has to be unique.
```go
func (c *Component) Type() string {
	return "example"
}
```

#### --- Component Methods ---

#### Component.Type()
Returns component type

```go
compType := component.Type()
```

### System
System manipulates entities and their components data.

#### --- System Interface ---

```go
type System interface {
	GetType() string
	GetPriority() int
	GetFrequency() uint
	Setup(entityStore *EntityStore)
	Process(entityStore *EntityStore, dt time.Duration)
	Cleanup(entityStore *EntityStore)
}
```

#### --- System Methods ---

#### System.GetType()
Returns system type, it has to be unique

```go
func (sys *ExampleSystem) GetType() string {
	// define system type here
	return "sys_example"
}

sysType := system.GetType()
```

#### System.GetPriority()
Returns system priority, higher number means higher priority

```go
func (sys *ExampleSystem) GetPriority() int {
	// define system priority here
	return 0
}
```

#### System.GetFrequency()
Returns process frequency in milliseconds

```go
func (sys *ExampleSystem) GetFrequency() uint {
	// define system frequency here
	// return zero (0) to call as often as possible
	return 0
}
```

#### System.Setup(entityStore *EntityStore)
Setup is called once before the main loop

```go
// use ECS_CORE to setup systems
ecs.Setup()

// main loop
for true {
	ecs.Process()
}
```

#### System.Process(entityStore *EntityStore, dt time.Duration)
Runs system logic once, has to be called in the main loop

```go
func (s *MovementSystem) Process(es *core.EntityStore, dt time.Duration) {
	finder := core.MakeFinder(es)
	entities := finder.Has("position").GetMany()

	for _, e := range entities {
		comp := *e.GetOne("position")
		pos := comp.(*PositionComponent)

		pos.X 
		pos.Y += 2
	}
}
```

#### System.Cleanup(entityStore *EntityStore)
Cleanup is called once after the main loop

```go
needToExit := true

// main loop
for !needToExit {
	ecs.Process()
}

// use ECS_CORE
ecs.Cleanup()
```

### ECS Core
Core links systems and entities, it provides process loop and setup/cleanup methods.

#### --- ECS Constructor ---
```go
ecs := core.MakeECS()
```

#### --- ECS Methods ---

#### ECS.Setup()
Runs all systems setup once

```go
ecs.Setup()
```

#### ECS.Process()
Runs all systems process once

```go
ecs.Process()
```

#### ECS.Cleanup()
Runs all systems cleanup once

### Entity Store
Stores entities and provides convenient management methods.

#### --- Entity Store Constructor ---
```go
es := core.MakeEntityStore()
```

#### --- Entity Store Methods ---

#### EntityStore.New(components ...Component) Entity
Creates a new entity with provided components

```go
posComp := MakePositionComponent(0, 0)
velComp := MakeVelocityComponent(1, 1)
entity := ecs.EntityStore.New(posComp, velComp)
```

#### EntityStore.GetComponentsById(entityId EntityID) []Component
Returns all components by entity id

```go
comps := ecs.EntityStore.GetComponentsById(1)
```

#### EntityStore.AddTo(entity Entity, components ...Component)
Adds components to an entity by ID

```go
ecs.EntityStore.AddTo(1, posComp, velComp)
```

#### EntityStore.RemoveFrom(entity Entity, componentTypes ...string)
Removes components from an entity by ID

```go
ecs.EntityStore.RemoveFrom(1, "position", "velocity")
```

#### EntityStore.Remove(entity Entity)
Removes an entity by id

```go
ecs.EntityStore.Remove(1)
```

#### EntityStore.GetAll() []*Entity
Returns a list of all entities

```go
entities := ecs.EntityStore.GetAll()
```

### System Store
Stores systems and provides convenient management methods.

#### --- System Store Constructor ---
```go
ss := core.MakeSystemStore()
```

#### --- System Store Methods ---

#### SystemStore.Add(system System)
Adds system to the store

```go
ecs.SystemStore.Add(sys)
```

#### SystemStore.Remove(typeName string)
Removes system from the store

```go
ecs.SystemStore.Remove("sys_example")
```

#### SystemStore.Get(typeName string) System
Returns system by type (string)

```go
sys := ecs.SystemStore.Get("sys_example")
```

#### SystemStore.GetAll() []System
Returns a list of all systems

```go
systems := ecs.SystemStore.GetAll()
```

#### SystemStore.GetPriority() []_SystemPriority
Returns a list of all systems sorted by priority

```go
priorities := ecs.SystemStore.GetPriority()

mostPrioritySys := priorities[0].GetType() // string
lastPrioritySys := priorities[len(systems)-1].GetType() // string
```

#### SystemStore.GetLastCallTime() map[string]time.Time
Returns a map of all systems last call time

```go
lastCallMap := ecs.SystemStore.GetLastCallTime()

moveSysLastCall := lastCallMap["sys_example"]
```