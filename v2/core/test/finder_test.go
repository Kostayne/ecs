package engine_test

import (
	"testing"

	. "github.com/kostayne/ecs/v2/core"
)

func TestMakeFinder(t *testing.T) {
	t.Run("Finder should return all entities if no filter are applied", func(t *testing.T) {
		expected := 3
		es := MakeEntityStore()

		es.New()
		es.New()
		es.New()

		f := MakeFinder(es)
		tg := len(f.GetMany())

		if len(f.GetMany()) != tg {
			t.Errorf("Expected %d entities, got %d", expected, tg)
		}
	})
}

func TestFinderWhere(t *testing.T) {
	t.Run("Finder.Where() should return empty list if 0 entities matched", func(t *testing.T) {
		expected := 0
		es := MakeEntityStore()

		es.New()
		es.New()
		es.New()

		f := MakeFinder(es)

		f.Where(func(e Entity) bool {
			return e.Id() > 500
		})

		tg := len(f.GetMany())

		if tg != expected {
			t.Errorf("Expected %d entities, got %d", expected, tg)
		}
	})

	t.Run("Finder.Where() should return correct entities count", func(t *testing.T) {
		expected := 2
		es := MakeEntityStore()

		es.New(&_TestComponent{})
		es.New(&_TestComponent{})
		es.New()

		f := MakeFinder(es)

		f.Where(func(e Entity) bool {
			return e.Has("TestComponent")
		})

		tg := len(f.GetMany())

		if tg != expected {
			t.Errorf("Expected %d entities, got %d", expected, tg)
		}
	})
}

func TestFinderHas(t *testing.T) {
	t.Run("Finder.Has() should return empty list if 0 entities matched", func(t *testing.T) {
		es := MakeEntityStore()
		es.New(&_TestComponent2{})
		es.New(&_TestComponent2{})
		es.New(&_TestComponent2{})

		f := MakeFinder(es)

		f.Has("UnknownComponent")
		tg := len(f.GetMany())

		if tg != 0 {
			t.Errorf("Expected 0 entities, got %d", tg)
		}
	})

	t.Run("Finder.Has should return correct entities count", func(t *testing.T) {
		expected := 2

		es := MakeEntityStore()

		es.New(&_TestComponent{})
		es.New(&_TestComponent{})
		es.New(&_TestComponent2{})

		f := MakeFinder(es)
		f.Has("TestComponent")

		tg := len(f.GetMany())

		if tg != expected {
			t.Errorf("Expected %d entities, got %d", expected, tg)
		}
	})
}
