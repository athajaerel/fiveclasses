package tbd

import (
	"testing"
)

func TestActionInterface(t *testing.T) {
//	// verify (?) that Buff implements State
//	var _ State = (*Buff)(nil)
}

// Actions may interact with Entities, Spaces and other Actions (triggering)
// Not with States or Forms

func TestActionEntityRelationship(t *testing.T) {
	// An Entity can have a Action
	// An Action can have an Entity
}

func TestActionActionRelationship(t *testing.T) {
	// An Action can have an Action
	// Triggering. OTOH, is this needed?
}

// see action_test.go for action-space relationship test

