package core

import "github.com/kostayne/ecs/utils"

// Finds entities by criteria
type FinderI interface {
	Get() Entity
	GetMany() []Entity

	Has(components ...string) *FinderI
	Where(predicate func(*Entity) bool) *FinderI
}

type Finder struct {
	es        *EntityStore
	entityIds []EntityID

	FinderI
}

func MakeFinder(es *EntityStore) *Finder {
	ids := make([]EntityID, 0)

	for id := range es.entities {
		ids = append(ids, id)
	}

	return &Finder{
		es:        es,
		entityIds: ids,
	}
}

// Filters entities by components presence.
func (f *Finder) Has(components ...string) *Finder {
	for _, c := range components {
		for i, id := range f.entityIds {
			_, isCompExists := f.es.ec_map[id][c]

			if !isCompExists {
				f.entityIds = utils.FastRemoveI(f.entityIds, i)
			}
		}
	}

	return f
}

// Filters entities by provided predicate.
func (f *Finder) Where(predicate func(*Entity) bool) *Finder {
	if predicate == nil {
		return f
	}

	for i, id := range f.entityIds {
		e := f.es.entities[id]

		if !predicate(&e) {
			f.entityIds = utils.FastRemoveI(f.entityIds, i)
		}
	}

	return f
}

// Returns all matched entities list
func (f *Finder) GetMany() []Entity {
	entities := make([]Entity, len(f.entityIds))

	for i, id := range f.entityIds {
		entities[i] = makeEntity(id, f.es)
	}

	return entities
}

// Returns first matched entity
func (f *Finder) GetOne() Entity {
	if len(f.entityIds) == 0 {
		return nil
	}

	return makeEntity(f.entityIds[0], f.es)
}
