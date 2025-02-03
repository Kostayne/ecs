package core

import "time"

// Core of the ECS engine.
type ECS_Core interface {
	Setup()
	Process()
	Cleanup()
}

// Default ECS core implementation.
type ECS_Default struct {
	EntityStore EntityStore
	SystemStore SystemStore
}

// Creates a new ECS instance.
func MakeECS() *ECS_Default {
	return &ECS_Default{
		EntityStore: *MakeEntityStore(),
		SystemStore: *MakeSystemStore(),
	}
}

// Runs all systems Setup method considering their priority.
func (e *ECS_Default) Setup() {
	for _, s := range e.SystemStore.GetAll() {
		s.Setup(&e.EntityStore)
	}
}

// Runs all systems Process method considering their frequency and priority.
func (e *ECS_Default) Process() {
	now := time.Now()
	systems := e.SystemStore.GetAll()
	callTime := e.SystemStore.GetLastCallTime()

	for _, p := range e.SystemStore.GetPriority() {
		s := systems[p.system]
		elapsed := now.Sub(callTime[p.system])

		if elapsed >= (time.Duration(s.GetFrequency()) * time.Millisecond) {
			s.Process(&e.EntityStore, elapsed)
			callTime[p.system] = now
		}
	}
}

// Runs all systems Cleanup method considering their priority.
func (e *ECS_Default) Cleanup() {
	for _, s := range e.SystemStore.GetAll() {
		s.Cleanup(&e.EntityStore)
	}
}
