package engine_test

import (
	"testing"

	. "github.com/kostayne/ecs/core"
)

type _MovementSystem struct {
	System
}

func (s *_MovementSystem) GetType() string {
	return "sys_movement"
}

func TestMakeSystemStore(t *testing.T) {
	ss := MakeSystemStore()

	if len(ss.Systems) != 0 {
		t.Errorf("Expected empty Systems slice, got %d elements", len(ss.Systems))
	}
}

func TestSystemStoreAdd(t *testing.T) {
	ss := MakeSystemStore()
	system := &_MovementSystem{}

	ss.Add(system)

	t.Run("Unique system should be added", func(t *testing.T) {
		if len(ss.Systems) != 1 {
			t.Errorf("Expected 1 system, got %d", len(ss.Systems))
		}
	})

	t.Run("Non-unique system should not be added", func(t *testing.T) {
		ss.Add(system)

		if len(ss.Systems) != 1 {
			t.Errorf("Expected 1 system, got %d", len(ss.Systems))
		}
	})
}

func TestSystemStoreRemove(t *testing.T) {
	ss := MakeSystemStore()
	system1 := &_MovementSystem{}

	ss.Add(system1)

	t.Run("Non-existent system should not be removed", func(t *testing.T) {
		ss.Remove("non_existent_type")

		if len(ss.Systems) != 1 {
			t.Errorf("Expected 1 system, got %d", len(ss.Systems))
		}
	})

	t.Run("Existing system should be removed", func(t *testing.T) {
		ss.Remove("sys_movement")

		if len(ss.Systems) != 0 {
			t.Errorf("Expected 0 systems, got %d", len(ss.Systems))
		}
	})
}
