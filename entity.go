package tbd

import ()

type Entity interface {
	listStates() []State
	addState(State)
	removeState(string)
	SetParent(* Space)
}

// TODO: move helpers to neutral helper unit?
func addStateHelper(states []State, ns State) []State {
	for _, s := range states {
		if s.Index() == ns.Index() {
			return states
		}
	}
	return append(states, ns)
}

func removeStateHelper(states []State, str string) []State {
	var tempVal []State
	for _, s := range states {
		if str != s.Index() {
			tempVal = append(tempVal, s)
		}
	}
	return tempVal
}

