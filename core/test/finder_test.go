package engine_test

import (
	"testing"

	. "github.com/kostayne/ecs/core"
)

func TestMakeFinder(t *testing.T) {
	t.Run("Finder should return all entities if no filter are applied", func(t *testing.T) {
		es := MakeEntityStore()
		es.AddNew()
		es.AddNew()
		es.AddNew()

		f := MakeFinder(es)

		if len(f.GetMany()) != len(es.Entities) {
			t.Errorf("Expected %d entities, got %d", len(es.Entities), len(f.GetMany()))
		}
	})
}

func TestFinderWhere(t *testing.T) {
	es := MakeEntityStore()

	es.AddNew(&_TestComponent{})
	es.AddNew(&_TestComponent{})
	es.AddNew()

	f := MakeFinder(es)

	f.Where(func(e *Entity) bool {
		return e.Has("TestComponent")
	})

	if len(f.GetMany()) != 2 {
		t.Errorf("Expected 2 entities, got %d", len(f.GetMany()))
	}
}

func TestFinderHas(t *testing.T) {
	es := MakeEntityStore()

	es.AddNew(&_TestComponent{})
	es.AddNew(&_TestComponent{})
	es.AddNew()

	f := MakeFinder(es)

	f.Has("TestComponent")

	if len(f.GetMany()) != 2 {
		t.Errorf("Expected 2 entities, got %d", len(f.GetMany()))
	}
}
