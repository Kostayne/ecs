package core

type System interface {
	GetType() string
	Setup(entityStore *EntityStore)
	Process(entityStore *EntityStore)
	Cleanup(entityStore *EntityStore)
}
