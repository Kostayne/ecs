package engine_test

import (
	"testing"

	. "github.com/kostayne/ecs/v2/core"
)

func TestMakeFinder(t *testing.T) {
	t.Run("Finder should return all entities if no filter are applied", func(t *testing.T) {
		es := MakeEntityStore()

		es.New()
		es.New()
		es.New()

		f := MakeFinder(es)

		if len(f.GetMany()) != len(es.GetAll()) {
			t.Errorf("Expected %d entities, got %d", len(es.GetAll()), len(f.GetMany()))
		}
	})
}

func TestFinderWhere(t *testing.T) {
	es := MakeEntityStore()

	es.New(&_TestComponent{})
	es.New(&_TestComponent{})
	es.New()

	f := MakeFinder(es)

	f.Where(func(e *Entity) bool {
		return (*e).Has("TestComponent")
	})

	if len(f.GetMany()) != 2 {
		t.Errorf("Expected 2 entities, got %d", len(f.GetMany()))
	}
}

func TestFinderHas(t *testing.T) {
	es := MakeEntityStore()

	es.New(&_TestComponent{})
	es.New(&_TestComponent{})
	es.New()

	f := MakeFinder(es)

	f.Has("TestComponent")

	if len(f.GetMany()) != 2 {
		t.Errorf("Expected 2 entities, got %d", len(f.GetMany()))
	}
}
