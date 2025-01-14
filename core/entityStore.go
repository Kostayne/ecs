package core

import "github.com/kostayne/ecs/utils"

type EntityStore struct {
	maxId    EntityID
	Entities []*Entity
}

func MakeEntityStore() *EntityStore {
	return &EntityStore{
		maxId:    0,
		Entities: make([]*Entity, 0),
	}
}

func (es *EntityStore) AddNew(components ...Component) *Entity {
	e := MakeEntity(es.maxId, components...)

	es.Entities = append(es.Entities, e)
	es.maxId++

	return e
}

func (es *EntityStore) Remove(e *Entity) {
	es.Entities = utils.FastRemove(es.Entities, e)
}

func (es *EntityStore) GetById(id EntityID) *Entity {
	for _, e := range es.Entities {
		if e.Id() == id {
			return e
		}
	}

	return nil
}
