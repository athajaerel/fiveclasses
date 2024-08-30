package tbd

import (
	"testing"
)

func TestSpaceInterface(t *testing.T) {
//	// verify (?) that Buff implements State
//	var _ State = (*Buff)(nil)
}

// Spaces may interact with Entities, States and Actions
// Not with Forms or other Spaces

func TestSpaceEntityRelationship(t *testing.T) {
	// A Space can have Entities
	// An Entity can have a Space
}

func TestSpaceActionRelationship(t *testing.T) {
	// A Space can have Actions
	// An Action can refer to Spaces
}

// see state_test.go for space-state relationship test

