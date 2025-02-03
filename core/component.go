package core

// Component is just an interface. Define your data in custom implementation struct.
type Component interface {
	// Returns component type (string).
	Type() string
}
