package core

type SystemStore struct {
	systems map[string]System
}

func MakeSystemStore() *SystemStore {
	return &SystemStore{
		systems: make(map[string]System),
	}
}

func (ss *SystemStore) Add(system System) {
	ss.systems[system.GetType()] = system
}

func (ss *SystemStore) Remove(typeName string) {
	delete(ss.systems, typeName)
}

func (ss *SystemStore) Get(typeName string) System {
	return ss.systems[typeName]
}

func (ss *SystemStore) GetAll() map[string]System {
	return ss.systems
}
