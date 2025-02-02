package engine_test

import (
	. "github.com/kostayne/ecs/v2/core"
)

type _TestComponent struct{}

func (t *_TestComponent) Type() string {
	return "TestComponent"
}

type _TestComponent2 struct {
	Component
}

func (t *_TestComponent2) Type() string {
	return "TestComponent2"
}
