package tbd

import ()

type Space interface {
	listStates() []State
	addState(State)
	removeState(string)
	IsEmpty() bool
	Put(Entity) // should store a backref? yes
	Count() uint
}

