package core

// Observes component attachments and detachments to notify related systems.
type Observer interface {
	// Returns observed component types.
	GetObservedTypes() []string
	// Sets observed component types.
	SetObservedTypes(types ...string)
	// Notifies systems about component attachment.
	OnAttach(componentType string, e Entity)
	// Notifies systems about component detachment.
	OnDetach(componentType string, e Entity)
}

// Base observer implementation for notifying systems.
type BaseObserver struct {
	systemStore     *SystemStore
	observedTypes   []string
	systemsToNotify []string
}

// Sets systems to be notified.
func (o *BaseObserver) SetNotifiableSystems(systems ...string) {
	o.systemsToNotify = systems
}

// Returns notifiable systems.
func (o *BaseObserver) GetNotifiableSystems() []string {
	return o.systemsToNotify
}

// Returns observed component types.
func (o *BaseObserver) GetObservedTypes() []string {
	return o.observedTypes
}

// Sets observed component types.
func (o *BaseObserver) SetObservedTypes(types ...string) {
	o.observedTypes = types
}

// Notifies systems with observable hooks about component attachment.
func (o *BaseObserver) OnAttach(componentType string, e Entity) {
	for _, s := range o.systemsToNotify {
		sys := o.systemStore.Get(s)

		if sys == nil {
			continue
		}

		if sys, ok := sys.(SystemWithObservableHooks); ok {
			sys.OnComponentAttached(componentType, e)
		}
	}
}

// Notifies systems with observable hooks about component detachment.
func (o *BaseObserver) OnDetach(componentType string, e Entity) {
	for _, s := range o.systemsToNotify {
		sys := o.systemStore.Get(s)

		if sys == nil {
			continue
		}

		if sys, ok := sys.(SystemWithObservableHooks); ok {
			sys.OnComponentDetached(componentType, e)
		}
	}
}

// Base observer constructor.
func NewObserver(systemStore *SystemStore) *BaseObserver {
	return &BaseObserver{
		systemStore:     systemStore,
		observedTypes:   []string{},
		systemsToNotify: []string{},
	}
}
