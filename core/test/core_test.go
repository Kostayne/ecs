package engine_test

import (
	"testing"
	"time"

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

func (s *_TestSystem) GetType() string    { return "sys_test" }
func (s *_TestSystem) GetPriority() int   { return 0 }
func (s *_TestSystem) GetFrequency() uint { return 0 }

func (s *_TestSystem) Setup(entityStore *EntityStore) {
	s.IsSetupCalled = true
}

func (s *_TestSystem) Process(entityStore *EntityStore) {
	s.IsProcessCalled = true
}

func (s *_TestSystem) Cleanup(entityStore *EntityStore) {
	s.IsCleanupCalled = true
}

func TestSystemHooks(t *testing.T) {
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

func TestSystemsPriority(t *testing.T) {
	t.Run("Systems should be processed in priority order", func(t *testing.T) {
		ecs := MakeECS()

		var prevCallIndex int8 = -1
		sysA := make_TEST_CORE_SYS_A(&prevCallIndex)
		sysB := make_TEST_CORE_SYS_B(&prevCallIndex)

		ecs.SystemStore.Add(sysA)
		ecs.SystemStore.Add(sysB)

		time.Sleep(time.Duration(sysB.GetFrequency()) * time.Millisecond)

		ecs.Process()

		if sysA.CallIndex != 0 || sysB.CallIndex != 1 {
			t.Errorf(
				"Expected a call index of 0 and b call index of 1, got %d and %d",
				sysA.CallIndex,
				sysB.CallIndex,
			)
		}
	})
}

func TestSystemsFrequency(t *testing.T) {
	t.Run("Systems should be processed after frequency time elapsed", func(t *testing.T) {
		ecs := MakeECS()

		var prevCallIndex int8 = -1
		sysA := make_TEST_CORE_SYS_A(&prevCallIndex)
		sysB := make_TEST_CORE_SYS_B(&prevCallIndex)

		ecs.SystemStore.Add(sysA)
		ecs.SystemStore.Add(sysB)

		ecs.Process()

		if sysB.CalledTimes != 0 {
			t.Errorf("Expected system B to not be called, got %d", sysB.CalledTimes)
		}

		time.Sleep(time.Millisecond * time.Duration(sysB.GetFrequency()))

		ecs.Process()

		if sysA.CalledTimes != 2 {
			t.Errorf("Expected system A to be called twice, got %d", sysA.CalledTimes)
		}

		if sysB.CalledTimes != 1 {
			t.Errorf("Expected system B to be called once, got %d", sysB.CalledTimes)
		}
	})
}
