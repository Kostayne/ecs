package core

type ComponentType = string

// Component to Entity map, needed for lookup by component type.
type CE_Map = map[ComponentType]map[EntityID]Component

// Entity to Component map, needed for lookup by entity ID.
type EC_Map = map[EntityID]map[ComponentType]Component

// Internal, needed for storing empty entities.
type _EntityMap = map[EntityID]Entity

// Stores entities and provides convenient management methods.
type EntityStore struct {
	maxId EntityID

	ce_map CE_Map
	ec_map EC_Map

	entities map[EntityID]Entity
}

// Entity store constructor.
func MakeEntityStore() *EntityStore {
	return &EntityStore{
		maxId: 0,

		ce_map: make(CE_Map),
		ec_map: make(EC_Map),

		entities: make(_EntityMap),
	}
}

// Creates a new entity and attaches provided components to it.
func (es *EntityStore) New(components ...Component) Entity {

	for _, c := range components {
		es.AddTo(es.maxId, c)
	}

	e := makeEntity(es.maxId, es)
	es.entities[es.maxId] = e

	es.maxId++
	return e
}

// Removes an entity by entity id, also detaches all components from the entity.
func (es *EntityStore) Remove(id EntityID) {
	for cType := range es.ec_map[id] {
		delete(es.ce_map[cType], id)
		delete(es.ec_map[id], cType)
	}

	delete(es.entities, id)
}

// Attaches components to an entity by ID.
func (es *EntityStore) AddTo(id EntityID, components ...Component) {
	for _, c := range components {
		cType := c.Type()

		if es.ce_map[cType] == nil {
			es.ce_map[cType] = make(map[EntityID]Component)
		}

		if es.ec_map[id] == nil {
			es.ec_map[id] = make(map[ComponentType]Component)
		}

		es.ce_map[cType][id] = c
		es.ec_map[id][cType] = c
	}
}

// Detaches components from an entity by entity ID.
func (es *EntityStore) RemoveFrom(id EntityID, componentTypes ...string) {
	for _, cType := range componentTypes {
		delete(es.ce_map[cType], id)
		delete(es.ec_map[id], cType)
	}
}

// Returns a list of attached components by entity ID.
func (es *EntityStore) GetById(id EntityID) []Component {
	e := es.entities[id]

	if e == nil {
		return nil
	}

	return e.GetAll()
}

// Returns a list of all stored entities.
func (es *EntityStore) GetAll() []Entity {
	entities := make([]Entity, len(es.ec_map))

	for _, e := range es.entities {
		entities = append(entities, e)
	}

	return entities
}
