package engine_test

import (
	"testing"

	. "github.com/kostayne/ecs/core"
)

type _FnMock struct {
	IsCalled bool
}

type _TestSystem struct {
	IsSetupCalled   bool
	IsProcessCalled bool
	IsCleanupCalled bool
}

func (s *_TestSystem) GetType() string { return "sys_test" }

func (s *_TestSystem) Setup(entityStore *EntityStore) {
	s.IsSetupCalled = true
}

func (s *_TestSystem) Process(entityStore *EntityStore) {
	s.IsProcessCalled = true
}

func (s *_TestSystem) Cleanup(entityStore *EntityStore) {
	s.IsCleanupCalled = true
}

func TestEngine(t *testing.T) {
	engine := MakeECS()

	sys := &_TestSystem{}
	engine.SystemStore.Add(sys)

	t.Run("Setup should call system.Setup", func(t *testing.T) {
		engine.Setup()

		if !sys.IsSetupCalled {
			t.Errorf("Expected system.Setup to be called")
		}
	})

	t.Run("Process should call system.Process", func(t *testing.T) {
		engine.Process()

		if !sys.IsProcessCalled {
			t.Errorf("Expected system.Process to be called")
		}
	})

	t.Run("Cleanup should call system.Cleanup", func(t *testing.T) {
		engine.Cleanup()

		if !sys.IsCleanupCalled {
			t.Errorf("Expected system.Cleanup to be called")
		}
	})

	engine.Setup()
}
