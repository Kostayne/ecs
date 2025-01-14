package core

import "github.com/kostayne/ecs/utils"

// Finds entities by criteria
type FinderI interface {
	GetOne() *Entity
	GetMany() []*Entity

	Has(components ...string) *Finder
	Where(predicate func(*Entity) bool) *Finder
}

type Finder struct {
	entities []*Entity

	FinderI
}

func MakeFinder(es *EntityStore) *Finder {
	return &Finder{
		entities: es.Entities,
	}
}

func (f *Finder) Has(components ...string) *Finder {
	for i, e := range f.entities {
		if !e.Has(components...) {
			f.entities = utils.FastRemoveI(f.entities, i)
		}
	}

	return f
}

func (f *Finder) Where(predicate func(*Entity) bool) *Finder {
	for i, e := range f.entities {
		if !predicate(e) {
			f.entities = utils.FastRemoveI(f.entities, i)
		}
	}

	return f
}

func (f *Finder) GetMany() []*Entity {
	return f.entities
}

func (f *Finder) GetOne() *Entity {
	return f.entities[0]
}
