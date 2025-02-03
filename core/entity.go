package core

type EntityID uint64

// Internal entity implementation struct, should not be used directly.
type DefaultEntity struct {
	id EntityID
	es *EntityStore
}

// Entity is just an interface, all data is stored in EntityStore as maps for performance reasons.
type Entity interface {
	// Returns entity ID (uint64).
	Id() EntityID

	// Returns true if entity has all provided components types.
	Has(componentTypes ...string) bool

	// Returns a component with provided type attached to the entity, may return nil if no such component exists.
	GetOne(componentType string) *Component

	// Returns a list of components attached to the entity with provided types.
	GetList(componentTypes ...string) []Component

	// Returns a list of all components attached to the entity.
	GetAll() []Component

	// Attaches provided components to the entity.
	Add(components ...Component)

	// Detaches provided component types from the entity.
	Remove(componentTypes ...string)
}

// A handy internal constructor.
func makeEntity(id EntityID, es *EntityStore) Entity {
	e := DefaultEntity{
		id: id,
		es: es,
	}

	return &e
}

// Returns entity ID (uint64).
func (e *DefaultEntity) Id() EntityID {
	return e.id
}

// Returns true if entity has all provided components types attached to it.
func (e *DefaultEntity) Has(componentTypes ...string) bool {
	comps := e.GetList(componentTypes...)

	return len(comps) == len(componentTypes)
}

// Returns an attached component by provided type, may return nil if no such component exists.
func (e *DefaultEntity) GetOne(componentType string) *Component {
	for _, c := range e.es.ec_map[e.id] {
		if c.Type() == componentType {
			return &c
		}
	}

	return nil
}

// Returns a list of components attached to the entity with provided types.
func (e *DefaultEntity) GetList(componentTypes ...string) []Component {
	comps := make([]Component, 0)

	for _, c := range e.es.ec_map[e.id] {
		for _, ct := range componentTypes {
			if c.Type() == ct {
				comps = append(comps, c)
			}
		}
	}

	return comps
}

// Returns a list of all components attached to the entity.
func (e *DefaultEntity) GetAll() []Component {
	comps := make([]Component, 0)

	for _, c := range e.es.ec_map[e.id] {
		comps = append(comps, c)
	}

	return comps
}

// Attaches provided components to the entity.
func (e *DefaultEntity) Add(components ...Component) {
	e.es.AddTo(e.id, components...)
}

// Detaches provided component types from the entity.
func (e *DefaultEntity) Remove(componentTypes ...string) {
	e.es.RemoveFrom(e.id, componentTypes...)
}
