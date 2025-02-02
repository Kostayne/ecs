package engine_test

import (
	"time"

	. "github.com/kostayne/ecs/v2/core"
)

// --- SYSTEM A
type _TEST_CORE_SYS_A struct {
	CallIndex   int8
	CalledTimes int16

	PrevCallIndex *int8
}

func (t *_TEST_CORE_SYS_A) Type() string {
	return "sys_a"
}

func (t *_TEST_CORE_SYS_A) Priority() int   { return 0 }
func (t *_TEST_CORE_SYS_A) Frequency() uint { return 0 }

func (t *_TEST_CORE_SYS_A) Setup(entityStore *EntityStore)   {}
func (t *_TEST_CORE_SYS_A) Cleanup(entityStore *EntityStore) {}

func (t *_TEST_CORE_SYS_A) Process(entityStore *EntityStore, dt time.Duration) {
	t.CallIndex = *t.PrevCallIndex + 1
	*t.PrevCallIndex = t.CallIndex

	t.CalledTimes++
}

// --- SYSTEM B
type _TEST_CORE_SYS_B struct {
	CallIndex   int8
	CalledTimes int16

	PrevCallIndex *int8
}

func (t *_TEST_CORE_SYS_B) Type() string {
	return "sys_b"
}

func (t *_TEST_CORE_SYS_B) Priority() int   { return -1 }
func (t *_TEST_CORE_SYS_B) Frequency() uint { return 15 }

func (t *_TEST_CORE_SYS_B) Setup(entityStore *EntityStore)   {}
func (t *_TEST_CORE_SYS_B) Cleanup(entityStore *EntityStore) {}

func (t *_TEST_CORE_SYS_B) Process(entityStore *EntityStore, dt time.Duration) {
	t.CallIndex = *t.PrevCallIndex + 1
	*t.PrevCallIndex = t.CallIndex

	t.CalledTimes++
}

// --- Constructors
func make_TEST_CORE_SYS_A(prevCallIndex *int8) *_TEST_CORE_SYS_A {
	return &_TEST_CORE_SYS_A{
		CallIndex:   0,
		CalledTimes: 0,

		PrevCallIndex: prevCallIndex,
	}
}

func make_TEST_CORE_SYS_B(prevCallIndex *int8) *_TEST_CORE_SYS_B {
	return &_TEST_CORE_SYS_B{
		CallIndex:   0,
		CalledTimes: 0,

		PrevCallIndex: prevCallIndex,
	}
}
