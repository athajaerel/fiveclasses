package tbd

import (
	"testing"
)

type World struct {
}

//type ? struct {
//}

func TestEntityInterface(t *testing.T) {
	// verify that BaseState implements State
//	var _ Entity = BaseEntity{}
//	var _ Entity = (*BaseEntity)(nil)
}

// Entities may interact with anything except other Entities
// see individual *_test.go for relationship tests

// make this a big test which tests all the things?
func TestWorld(t *testing.T) {
	// verify that World implements Entity
	var _ Entity = World{}
	var _ Entity = (*World)(nil)
	world := &World{}
}

