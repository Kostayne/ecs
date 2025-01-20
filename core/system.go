package core

type System interface {
	// Should return a unique system type string
	GetType() string

	// System priority, higher number means higher priority
	GetPriority() int

	// Process frequency in milliseconds
	GetFrequency() uint

	// Called once before the main loop
	Setup(entityStore *EntityStore)

	// Called in main loop
	Process(entityStore *EntityStore)

	// Called once after the main loop
	Cleanup(entityStore *EntityStore)
}
