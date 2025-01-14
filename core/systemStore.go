package core

import (
	"github.com/kostayne/ecs/utils"
)

type SystemStore struct {
	Systems []System
}

func MakeSystemStore() *SystemStore {
	return &SystemStore{
		Systems: make([]System, 0),
	}
}

func (ss *SystemStore) Add(system System) {
	for _, s := range ss.Systems {
		if s.GetType() == system.GetType() {
			println("System of type " + system.GetType() + " already exists!")
			return
		}
	}

	ss.Systems = append(ss.Systems, system)
}

func (ss *SystemStore) Remove(typeName string) {
	for i, s := range ss.Systems {
		if s.GetType() == typeName {
			ss.Systems = utils.ShiftRemoveI(ss.Systems, i)
			return
		}
	}

	println("System of type " + typeName + " not found!")
}
