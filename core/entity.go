package core

import (
	"github.com/kostayne/ecs/utils"
)

type EntityID uint64

type Entity struct {
	id         EntityID
	Components []Component `json:"components"`
}

// Get entity id
func (e *Entity) Id() EntityID {
	return e.id
}

// Add entity components
func (e *Entity) Add(components ...Component) {
	for _, c := range components {
		isUnique := true

		for _, ec := range e.Components {
			if ec.Type() == c.Type() {
				isUnique = false
				break
			}
		}

		if !isUnique {
			println("Entity already has a component of type " + c.Type())
			continue
		}

		e.Components = append(e.Components, c)
	}
}

// Remove entity components
func (e *Entity) Remove(components ...Component) {
	for _, c := range components {
		for i, ec := range e.Components {
			if ec.Type() == c.Type() {
				e.Components = utils.FastRemoveI(e.Components, i)
				break
			}
		}
	}
}

// Check if entity has components
func (e *Entity) Has(componentTypes ...string) bool {
	res := true

	for _, ct := range componentTypes {
		found := false

		for _, ec := range e.Components {
			if ec.Type() == ct {
				found = true
				break
			}
		}

		if !found {
			res = false
			break
		}
	}

	return res
}

// Get entity components list
func (e *Entity) GetList(componentTypes ...string) []*Component {
	res := make([]*Component, len(componentTypes))

	for i, ct := range componentTypes {
		for _, ec := range e.Components {
			if ec.Type() == ct {
				res[i] = &ec
				break
			}
		}
	}

	return res
}

// Get entity component
func (e *Entity) Get(componentType string) *Component {
	for _, ec := range e.Components {
		if ec.Type() == componentType {
			return &ec
		}
	}

	return nil
}

// Create new entity
func MakeEntity(id EntityID, components ...Component) *Entity {
	return &Entity{
		Components: components,
	}
}
