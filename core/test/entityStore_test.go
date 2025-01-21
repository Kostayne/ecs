package engine_test

import (
	"testing"

	. "github.com/kostayne/ecs/core"
)

func TestEntityStoreAddNew(t *testing.T) {
	es := MakeEntityStore()

	t.Run("Entity should be added", func(t *testing.T) {
		es.New()

		if len(es.GetAll()) != 1 {
			t.Errorf("Expected 1 entity, got %d", len(es.GetAll()))
		}
	})

	t.Run("First entity ID should be 0", func(t *testing.T) {
		e := es.New()

		if e.Id() == 0 {
			t.Errorf("Expected entity.ID to be 0, got %d", e.Id())
		}
	})

	t.Run("Entities id should increment", func(t *testing.T) {
		e1 := es.New()
		e2 := es.New()

		if (e2.Id() - e1.Id()) != 1 {
			t.Errorf("Expected consecutive ID difference to be 1, got %d", e2.Id()-e1.Id())
		}
	})
}

func TestEntityStoreRemove(t *testing.T) {
	es := MakeEntityStore()
	e := es.New()

	es.Remove(e.Id())

	if len(es.GetAll()) != 0 {
		t.Errorf("Expected Entities to be empty, got %d elements", len(es.GetAll()))
	}
}

func TestEntityStoreGetById(t *testing.T) {
	es := MakeEntityStore()

	t.Run("Non-existent entity should not be found", func(t *testing.T) {
		found := es.GetComponentsById(515)

		if found != nil {
			t.Errorf("Expected entity not to be found, got %v", found)
		}
	})

	t.Run("Existing entity should be found", func(t *testing.T) {
		e := es.New()
		found := es.GetComponentsById(e.Id())

		if found == nil {
			t.Errorf("Expected entity to be found, got nil")
		}
	})
}
