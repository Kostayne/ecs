package core

import "time"

// Core of the ECS engine.
type ECS struct {
	EntityStore EntityStore
	SystemStore SystemStore
}

// Creates a new ECS instance.
func MakeECS() *ECS {
	return &ECS{
		EntityStore: *MakeEntityStore(),
		SystemStore: *MakeSystemStore(),
	}
}

// Runs all systems Setup method considering their priority.
func (e *ECS) Setup() {
	for _, p := range e.SystemStore.priority {
		s := e.SystemStore.systems[p.system]
		s.Setup(&e.EntityStore)
	}
}

// Runs all systems Process method considering their frequency and priority.
func (e *ECS) Process() {
	now := time.Now()
	systems := e.SystemStore.GetAll()
	callTime := e.SystemStore.LastCallTimeMap()

	for _, p := range e.SystemStore.Priority() {
		s := systems[p.system]
		elapsed := now.Sub(callTime[p.system])

		if elapsed >= (time.Duration(s.Frequency()) * time.Millisecond) {
			s.Process(&e.EntityStore, elapsed)
			callTime[p.system] = now
		}
	}
}

// Runs all systems Cleanup method considering their priority.
func (e *ECS) Cleanup() {
	for _, p := range e.SystemStore.priority {
		s := e.SystemStore.systems[p.system]
		s.Cleanup(&e.EntityStore)
	}
}
