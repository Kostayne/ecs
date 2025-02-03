package core

import (
	"slices"
	"time"

	"github.com/kostayne/ecs/utils"
)

// Filters entities by attached to them components presence.
type _SystemPriority struct {
	priority int
	system   string
}

// Priority value getter, no setter allowed.
func (p *_SystemPriority) GetValue() int {
	return p.priority
}

// Priority value getter, no setter allowed.
func (p *_SystemPriority) GetSystemType() string {
	return p.system
}

// Manages all systems according to their priority & process frequency.
type SystemStore struct {
	systems map[string]System

	// Priority from high to low
	priority []_SystemPriority

	// Time from last process
	lastCallTime map[string]time.Time
}

// System store constructor.
func MakeSystemStore() *SystemStore {
	return &SystemStore{
		systems:      make(map[string]System),
		priority:     make([]_SystemPriority, 0),
		lastCallTime: make(map[string]time.Time),
	}
}

// Adds a system to the store, so it can be processed. Panics if the same system type is already added.
func (ss *SystemStore) Add(system System) {
	// --- Setting up the priority
	ss.priority = append(ss.priority, makeSystemPriority(system))

	// sort hight to low
	slices.SortFunc(ss.priority, func(a, b _SystemPriority) int {
		return b.priority - a.priority
	})

	// --- Adding the system
	ss.systems[system.GetType()] = system

	// Add the last call time
	ss.lastCallTime[system.GetType()] = time.Now()
}

// Removes a system from the store, so it can no longer be processed.
func (ss *SystemStore) Remove(typeName string) {
	// --- Removing the system
	delete(ss.systems, typeName)

	// --- Removing the priority
	for i, p := range ss.priority {
		if p.system == typeName {
			ss.priority = utils.ShiftRemoveI(ss.priority, i)
			break
		}
	}

	// Remove the last call time
	delete(ss.lastCallTime, typeName)
}

// Returns a system from the store by its type. May return nil if no such system was added.
func (ss *SystemStore) Get(typeName string) System {
	return ss.systems[typeName]
}

// Returns all systems from the store.
func (ss *SystemStore) GetAll() map[string]System {
	return ss.systems
}

// Returns all systems from the store sorted ascending by priority.
func (ss *SystemStore) GetPriority() []_SystemPriority {
	return ss.priority
}

// Returns map of last call times, key is system type, value is last call time. May be useful for debugging.
func (ss *SystemStore) GetLastCallTime() map[string]time.Time {
	return ss.lastCallTime
}

// Internal system priority constructor.
func makeSystemPriority(system System) _SystemPriority {
	return _SystemPriority{
		priority: system.GetPriority(),
		system:   system.GetType(),
	}
}
