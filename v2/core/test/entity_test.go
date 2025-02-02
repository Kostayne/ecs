package engine_test

import (
	"testing"

	. "github.com/kostayne/ecs/v2/core"
)

func TestEntityAdd(t *testing.T) {
	t.Run("Add single component", func(t *testing.T) {
		es := MakeEntityStore()
		e := es.New()

		e.Add(&_TestComponent{})

		if len(e.GetAll()) != 1 {
			t.Errorf("Expected 1 component, got %d", len(e.GetAll()))
		}
	})

	t.Run("Duplicate component", func(t *testing.T) {
		es := MakeEntityStore()
		e := es.New()

		e.Add(&_TestComponent{})
		e.Add(&_TestComponent{})

		if len(e.GetAll()) != 1 {
			t.Errorf("Expected 1 component, got %d", len(e.GetAll()))
		}
	})
}

func TestEntityRemove(t *testing.T) {
	t.Run("Remove single component", func(t *testing.T) {
		es := MakeEntityStore()
		e := es.New(&_TestComponent{})

		e.Remove("TestComponent")

		if len(e.GetAll()) != 0 {
			t.Errorf("Expected 0 components, got %d", len(e.GetAll()))
		}
	})

	t.Run("Remove non-existent component should not panic", func(t *testing.T) {
		es := MakeEntityStore()
		e := es.New()

		e.Remove("NonExistentComponent")
	})
}

func TestEntityHas(t *testing.T) {
	es := MakeEntityStore()
	e := es.New()

	t.Run("Has return false for non-existent component", func(t *testing.T) {
		if e.Has("TestComponent") {
			t.Errorf("Expected false, got true")
		}
	})

	t.Run("Has return true for existing component", func(t *testing.T) {
		e.Add(&_TestComponent{})

		if !e.Has("TestComponent") {
			t.Errorf("Expected true, got false")
		}
	})
}

func TestEntityGetList(t *testing.T) {
	es := MakeEntityStore()
	e := es.New()

	e.Add(&_TestComponent{})

	res := e.GetList("TestComponent")

	if len(res) != 1 {
		t.Errorf("Expected 1 component, got %d", len(res))
	}

	if res[0].Type() != "TestComponent" {
		t.Errorf("Expected TestComponent, got %s", res[0].Type())
	}
}

func TestEntityGet(t *testing.T) {
	es := MakeEntityStore()
	e := es.New()

	e.Add(&_TestComponent{})

	t.Run("Get return nil for non-existent component", func(t *testing.T) {
		res := e.GetOne("NonExistingComponent")

		if res != nil {
			t.Errorf("Expected nil, got %s", (*res).Type())
		}
	})

	t.Run("Get return existing component", func(t *testing.T) {
		res := e.GetOne("TestComponent")

		if res == nil {
			t.Errorf("Expected TestComponent, got nil")
		}

		if (*res).Type() != "TestComponent" {
			t.Errorf("Expected TestComponent, got %s", (*res).Type())
		}
	})
}
