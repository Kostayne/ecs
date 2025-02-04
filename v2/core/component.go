package core

// Component is just an interface. Define your data in custom implementation struct.
type Component interface {
	// Returns component type string
	Type() string
}

// Like Component, but with additional hooks (OnAttach, OnDetach). Use it when you need to access component owner.
type ComponentWithHooks interface {
	Component

	// Called when component is attached to an entity (added)
	OnAttach(owner Entity)

	// Called when component is detached from an entity (removed)
	OnDetach()
}
