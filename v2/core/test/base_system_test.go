package engine_test

import (
	"testing"
	"time"

	. "github.com/kostayne/ecs/v2/core"
)

type _TEST_BASE_SYS struct {
	*SystemBase
}

func (s *_TEST_BASE_SYS) Process(es *EntityStore, dt time.Duration) {}

const (
	sysBaseType     = "sys_base_test"
	sysBasePriority = -5
	sysBaseFreq     = 10
)

func makeTestBaseSys() *_TEST_BASE_SYS {
	return &_TEST_BASE_SYS{
		SystemBase: MakeSystemBase(sysBaseType, sysBaseFreq, sysBasePriority),
	}
}

func TestSystemBase(t *testing.T) {
	sys := makeTestBaseSys()

	t.Run("Type should match", func(t *testing.T) {
		if sys.Type() != sysBaseType {
			t.Errorf("Expected type to be 'sys_test', got '%s'", sys.Type())
		}
	})

	t.Run("Priority should match", func(t *testing.T) {
		if sys.Priority() != sysBasePriority {
			t.Errorf("Expected priority to be %d, got %d", sysBasePriority, sys.Priority())
		}
	})

	t.Run("Frequency should match", func(t *testing.T) {
		if sys.Frequency() != sysBaseFreq {
			t.Errorf("Expected frequency to be %d, got %d", sysBaseFreq, sys.Frequency())
		}
	})

	t.Run("Should be compatible with System interface", func(t *testing.T) {
		_ = System(sys)
	})
}
