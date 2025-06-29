package engine_test

import (
	"testing"
	"time"

	. "github.com/kostayne/ecs/v2/core"
)

type _ObserverTestComponent struct{}

func (c *_ObserverTestComponent) Type() string {
	return "observer_test_component"
}

type _ObserverTestSys struct {
	attachCalled []string
	detachCalled []string

	SystemBase
}

func (s *_ObserverTestSys) Process(es *EntityStore, dt time.Duration) {}

func newObserverTestSys() *_ObserverTestSys {
	return &_ObserverTestSys{
		attachCalled: make([]string, 0),
		detachCalled: make([]string, 0),
		SystemBase:   *MakeSystemBase("_observer_test_sys", 0, 0),
	}
}

func (s *_ObserverTestSys) OnComponentAttached(componentType string, e Entity) {
	s.attachCalled = append(s.attachCalled, componentType)
}

func (s *_ObserverTestSys) OnComponentDetached(componentType string, e Entity) {
	s.detachCalled = append(s.detachCalled, componentType)
}

func TestObservers(t *testing.T) {
	ecs := MakeECS()

	sys := newObserverTestSys()
	ecs.SystemStore.Add(sys)

	comp := &_ObserverTestComponent{}

	observer := NewObserver(&ecs.SystemStore)
	observer.SetNotifiableSystems(sys.Type())
	observer.SetObservedTypes(comp.Type())

	ecs.EntityStore.AddObserver(observer)

	t.Run("Attach should call system.OnComponentAttached", func(t *testing.T) {
		ecs.EntityStore.New(comp)

		tgSys := (ecs.SystemStore.Get(sys.Type())).(*_ObserverTestSys)

		if len(tgSys.attachCalled) == 0 {
			t.Errorf("Expected attached component notification to be called")
			return
		}

		if tgSys.attachCalled[0] != comp.Type() {
			t.Errorf("Expected attached component notification to be called with %s, got %s", comp.Type(), tgSys.attachCalled[0])
		}
	})

	t.Run("Detach should call system.OnComponentDetached", func(t *testing.T) {
		ents := ecs.EntityStore.GetAll()

		for _, ent := range ents {
			ecs.EntityStore.Remove(ent.Id())
		}

		tgSys := (ecs.SystemStore.Get(sys.Type())).(*_ObserverTestSys)

		if len(tgSys.detachCalled) == 0 {
			t.Errorf("Expected detached component notification to be called")
			return
		}

		if tgSys.detachCalled[0] != comp.Type() {
			t.Errorf("Expected detached component notification to be called with %s, got %s", comp.Type(), tgSys.detachCalled[0])
		}
	})

	t.Run("Attach & detach not panic if system is not registered", func(t *testing.T) {
		ecs.SystemStore.Remove(sys.Type())
		ecs.EntityStore.RemoveObserver(observer)

		ents := ecs.EntityStore.GetAll()

		for _, ent := range ents {
			ecs.EntityStore.Remove(ent.Id())
		}

		newObserver := NewObserver(&ecs.SystemStore)
		newObserver.SetNotifiableSystems("404_test_sys")
		newObserver.SetObservedTypes(comp.Type())

		ecs.EntityStore.AddObserver(newObserver)

		newEnt := ecs.EntityStore.New(comp)
		ecs.EntityStore.Remove(newEnt.Id())
	})
}
