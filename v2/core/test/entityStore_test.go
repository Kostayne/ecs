package engine_test

import (
	"testing"

	. "github.com/kostayne/ecs/v2/core"
)

type _ComponentWithHooks struct {
	OnAttachIsCalled bool
	OnDetachIsCalled bool
}

func (c *_ComponentWithHooks) Type() string {
	return "component_with_hooks"
}

func (c *_ComponentWithHooks) OnAttach(e Entity) {
	c.OnAttachIsCalled = true
}

func (c *_ComponentWithHooks) OnDetach() {
	c.OnDetachIsCalled = true
}

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

	t.Run("Component on Attached should be called on es.New()", func(t *testing.T) {
		c := &_ComponentWithHooks{}
		es.New(c)

		if !c.OnAttachIsCalled {
			t.Errorf("Expected OnAttach to be called")
		}
	})

	t.Run("Component on Attached should be called on es.AddTo()", func(t *testing.T) {
		e := es.New()
		c := &_ComponentWithHooks{}

		es.AddTo(e.Id(), c)

		if !c.OnAttachIsCalled {
			t.Errorf("Expected OnAttach to be called")
		}
	})

	t.Run("Component on Attached should be called on entity.Add()", func(t *testing.T) {
		e := es.New()
		c := &_ComponentWithHooks{}

		e.Add(c)

		if !c.OnAttachIsCalled {
			t.Errorf("Expected OnAttach to be called")
		}
	})

	t.Run("Component on Detached should be called on es.RemoveFrom()", func(t *testing.T) {
		e := es.New()
		c := &_ComponentWithHooks{}

		es.AddTo(e.Id(), c)
		es.RemoveFrom(e.Id(), c.Type())

		if !c.OnDetachIsCalled {
			t.Errorf("Expected OnDetach to be called")
		}
	})

	t.Run("Component on Detached should be called on entity.Remove()", func(t *testing.T) {
		e := es.New()
		c := &_ComponentWithHooks{}

		es.AddTo(e.Id(), c)
		e.Remove(c.Type())

		if !c.OnDetachIsCalled {
			t.Errorf("Expected OnDetach to be called")
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
		found := es.GetById(515)

		if found != nil {
			t.Errorf("Expected entity not to be found, got %v", found)
		}
	})

	t.Run("Existing entity should be found", func(t *testing.T) {
		e := es.New()
		found := es.GetById(e.Id())

		if found == nil {
			t.Errorf("Expected entity to be found, got nil")
		}
	})
}
