package core

import (
	"slices"
	"time"

	"github.com/kostayne/ecs/utils"
)

type _SystemPriority struct {
	priority int
	system   string
}

func (p *_SystemPriority) GetValue() int {
	return p.priority
}

func (p *_SystemPriority) GetSystemType() string {
	return p.system
}

type SystemStore struct {
	systems map[string]System

	// Priority from high to low
	priority []_SystemPriority

	// Time from last process
	lastCallTime map[string]time.Time
}

func MakeSystemStore() *SystemStore {
	return &SystemStore{
		systems:      make(map[string]System),
		priority:     make([]_SystemPriority, 0),
		lastCallTime: make(map[string]time.Time),
	}
}

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

func (ss *SystemStore) Get(typeName string) System {
	return ss.systems[typeName]
}

func (ss *SystemStore) GetAll() map[string]System {
	return ss.systems
}

func (ss *SystemStore) GetPriority() []_SystemPriority {
	return ss.priority
}

func (ss *SystemStore) GetLastCallTime() map[string]time.Time {
	return ss.lastCallTime
}

// --- Utils ---
func makeSystemPriority(system System) _SystemPriority {
	return _SystemPriority{
		priority: system.GetPriority(),
		system:   system.GetType(),
	}
}
