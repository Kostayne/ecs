package core

import "time"

type System interface {
	// Returns a unique system type (string)
	GetType() string

	// Returns system priority, higher number means higher priority
	GetPriority() int

	// Returns process frequency in milliseconds
	GetFrequency() uint

	// Called once before the main loop
	Setup(entityStore *EntityStore)

	// Called in main loop, dt is delta time (time since last call)
	Process(entityStore *EntityStore, dt time.Duration)

	// Called once after the main loop
	Cleanup(entityStore *EntityStore)
}
