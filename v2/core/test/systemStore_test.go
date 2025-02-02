package engine_test

import (
	"testing"
	"time"

	. "github.com/kostayne/ecs/v2/core"
)

type _MovementSystem struct {
	System
}

func (s *_MovementSystem) Type() string {
	return "sys_movement"
}

func (s *_MovementSystem) Priority() int   { return 0 }
func (s *_MovementSystem) Frequency() uint { return 0 }

func (s *_MovementSystem) Setup(es *EntityStore)                     {}
func (s *_MovementSystem) Process(es *EntityStore, dt time.Duration) {}
func (s *_MovementSystem) Cleanup(es *EntityStore)                   {}

func TestMakeSystemStore(t *testing.T) {
	ss := MakeSystemStore()

	if len(ss.GetAll()) != 0 {
		t.Errorf("Expected empty Systems slice, got %d elements", len(ss.GetAll()))
	}
}

func TestSystemStoreAdd(t *testing.T) {
	ss := MakeSystemStore()
	system := &_MovementSystem{}

	ss.Add(system)

	t.Run("Unique system should be added", func(t *testing.T) {
		if len(ss.GetAll()) != 1 {
			t.Errorf("Expected 1 system, got %d", len(ss.GetAll()))
		}
	})

	t.Run("Non-unique system should not be added (panic)", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic, got nil")
			}
		}()

		ss.Add(system)

		if len(ss.GetAll()) != 1 {
			t.Errorf("Expected 1 system, got %d", len(ss.GetAll()))
		}
	})
}

func TestSystemStoreRemove(t *testing.T) {
	ss := MakeSystemStore()
	system1 := &_MovementSystem{}

	ss.Add(system1)

	t.Run("Non-existent system should not be removed", func(t *testing.T) {
		ss.Remove("non_existent_type")

		if len(ss.GetAll()) != 1 {
			t.Errorf("Expected 1 system, got %d", len(ss.GetAll()))
		}
	})

	t.Run("Existing system should be removed", func(t *testing.T) {
		ss.Remove("sys_movement")

		if len(ss.GetAll()) != 0 {
			t.Errorf("Expected 0 systems, got %d", len(ss.GetAll()))
		}
	})
}
