package core

import "time"

type System interface {
	// Returns a unique system type (string)
	Type() string

	// Returns system priority, higher number means higher priority
	Priority() int

	// Returns process frequency in milliseconds
	Frequency() uint

	// Called once before the main loop
	Setup(entityStore *EntityStore)

	// Called in main loop, dt is delta time (time since last call)
	Process(entityStore *EntityStore, dt time.Duration)

	// Called once after the main loop
	Cleanup(entityStore *EntityStore)
}

// SystemBase implements the System interface but not includes Process.
// It can be used to reduce boilerplate code, override only the methods you need.
type SystemBase struct {
	frequency  uint
	priority   int
	systemType string
}

// Default implementation that eliminates boilerplate.
func (s *SystemBase) Type() string { return s.systemType }

// Default implementation that eliminates boilerplate.
func (s *SystemBase) Priority() int { return s.priority }

// Default implementation that eliminates boilerplate.
func (s *SystemBase) Frequency() uint { return s.frequency }

// Default implementation that does nothing.
func (s *SystemBase) Setup(entityStore *EntityStore) {}

// Default implementation that does nothing.
func (s *SystemBase) Cleanup(entityStore *EntityStore) {}

// Constructor for a SystemBase, stores system params. Use it to reduce boilerplate.
func MakeSystemBase(systemType string, frequency uint, priority int) *SystemBase {
	return &SystemBase{
		systemType: systemType,
		frequency:  frequency,
		priority:   priority,
	}
}
