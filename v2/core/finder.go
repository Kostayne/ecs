package core

type FinderI interface {
	Get() Entity
	GetMany() []Entity

	Has(components ...string) FinderI
	Where(predicate func(Entity) bool) FinderI
}

// Finder implementation, stores entity IDs matched by filters.
type Finder struct {
	es        *EntityStore
	entityIds []EntityID

	FinderI
}

// Default finder implementation constructor.
func MakeFinder(es *EntityStore) FinderI {
	ids := make([]EntityID, 0)

	for id := range es.entities {
		ids = append(ids, id)
	}

	return &Finder{
		es:        es,
		entityIds: ids,
	}
}

// Filters entities by attached to them components presence.
func (f *Finder) Has(components ...string) FinderI {
	matched := make([]EntityID, 0)

	for _, id := range f.entityIds {
		isMatching := true

		for _, c := range components {
			_, isCompExists := f.es.ec_map[id][c]

			if !isCompExists {
				isMatching = false
				break
			}
		}

		if isMatching {
			matched = append(matched, id)
		}
	}

	f.entityIds = matched
	return f
}

// Filters entities by provided predicate.
func (f *Finder) Where(predicate func(Entity) bool) FinderI {
	if predicate == nil {
		return f
	}

	matched := make([]EntityID, 0)

	for _, id := range f.entityIds {
		e := f.es.entities[id]

		if predicate(e) {
			matched = append(matched, id)
		}
	}

	f.entityIds = matched
	return f
}

// Returns all matched entities list.
func (f *Finder) GetMany() []Entity {
	entities := make([]Entity, len(f.entityIds))

	for i, id := range f.entityIds {
		entities[i] = makeEntity(id, f.es)
	}

	return entities
}

// Returns the first matched entity.
func (f *Finder) GetOne() Entity {
	if len(f.entityIds) == 0 {
		return nil
	}

	return makeEntity(f.entityIds[0], f.es)
}
