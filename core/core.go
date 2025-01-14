package core

type ECS_Core interface {
	Setup()
	Process()
	Cleanup()
}

type ECS_Default struct {
	EntityStore EntityStore
	SystemStore SystemStore

	ECS_Core
}

func MakeECS() *ECS_Default {
	return &ECS_Default{
		EntityStore: *MakeEntityStore(),
		SystemStore: *MakeSystemStore(),
	}
}

func (e *ECS_Default) Setup() {
	for _, s := range e.SystemStore.Systems {
		s.Setup(&e.EntityStore)
	}
}

func (e *ECS_Default) Process() {
	for _, s := range e.SystemStore.Systems {
		s.Process(&e.EntityStore)
	}
}

func (e *ECS_Default) Cleanup() {
	for _, s := range e.SystemStore.Systems {
		s.Cleanup(&e.EntityStore)
	}
}
