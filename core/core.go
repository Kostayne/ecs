package core

import "time"

type ECS_Core interface {
	Setup()
	Process()
	Cleanup()
}

type ECS_Default struct {
	EntityStore EntityStore
	SystemStore SystemStore
}

func MakeECS() *ECS_Default {
	return &ECS_Default{
		EntityStore: *MakeEntityStore(),
		SystemStore: *MakeSystemStore(),
	}
}

func (e *ECS_Default) Setup() {
	for _, s := range e.SystemStore.GetAll() {
		s.Setup(&e.EntityStore)
	}
}

func (e *ECS_Default) Process() {
	now := time.Now()
	systems := e.SystemStore.GetAll()
	callTime := e.SystemStore.GetLastCallTime()

	for _, p := range e.SystemStore.GetPriority() {
		s := systems[p.system]
		elapsed := now.Sub(callTime[p.system])

		if elapsed >= (time.Duration(s.GetFrequency()) * time.Millisecond) {
			s.Process(&e.EntityStore)
			callTime[p.system] = now
		}
	}
}

func (e *ECS_Default) Cleanup() {
	for _, s := range e.SystemStore.GetAll() {
		s.Cleanup(&e.EntityStore)
	}
}
