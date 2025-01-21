package core

type EntityID uint64

type DefaultEntity struct {
	id EntityID
	es *EntityStore
}

type Entity interface {
	// Returns entity ID
	Id() EntityID

	// Returns true if all components are present
	Has(componentTypes ...string) bool

	// Returns a specified component
	Get(componentType string) *Component

	// Returns a list of specified components
	GetList(componentTypes ...string) []Component
	GetAll() []Component

	// Adds provided components
	Add(components ...Component)

	// Removes components with provided types
	Remove(componentTypes ...string)
}

func makeEntity(id EntityID, es *EntityStore) Entity {
	e := DefaultEntity{
		id: id,
		es: es,
	}

	return &e
}

// Returns entity ID
func (e *DefaultEntity) Id() EntityID {
	return e.id
}

// Returns true if all components are present
func (e *DefaultEntity) Has(componentTypes ...string) bool {
	comps := e.GetList(componentTypes...)

	return len(comps) == len(componentTypes)
}

// Returns a specified component
func (e *DefaultEntity) Get(componentType string) *Component {
	for _, c := range e.es.ec_map[e.id] {
		if c.Type() == componentType {
			return &c
		}
	}

	return nil
}

// Returns a list of specified components
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

// Returns a list of all components
func (e *DefaultEntity) GetAll() []Component {
	comps := make([]Component, 0)

	for _, c := range e.es.ec_map[e.id] {
		comps = append(comps, c)
	}

	return comps
}

// Adds components
func (e *DefaultEntity) Add(components ...Component) {
	e.es.AddTo(e.id, components...)
}

// Removes components
func (e *DefaultEntity) Remove(componentTypes ...string) {
	e.es.RemoveFrom(e.id, componentTypes...)
}
