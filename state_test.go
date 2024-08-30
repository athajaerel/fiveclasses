package tbd

import (
	"testing"
)

type Marble struct {
	radius float32
	kind string
	states []State
	parentSpace *Space
}


func (be *Marble) addState(ns State) {
	be.states = addStateHelper(be.states, ns)
}

func (be *Marble) listStates() []State {
	return be.states
}

func (be *Marble) removeState(str string) {
	be.states = removeStateHelper(be.states, str)
}

func (be *Marble) SetParent(p *Space) {
	be.parentSpace = p
}

type Content struct {
	states []State
	// concrete Entities --- everything must be somewhere
	// use another way to storing pointers if needed
	entities []Entity
}

func (bp *Content) addState(ns State) {
	bp.states = addStateHelper(bp.states, ns)
}

func (bp *Content) listStates() []State {
	return bp.states
}

func (bp *Content) removeState(str string) {
	bp.states = removeStateHelper(bp.states, str)
}

func (bp *Content) IsEmpty() bool {
	return len(bp.entities) == 0
}

func (bp *Content) Put(e Entity) {
	bp.entities = append(bp.entities, e)
	var tmp *Space = bp
	e.SetParent(tmp)
}

func (bp *Content) Count() uint {
	return uint(len(bp.entities))
}

type Bag struct {
	states []State
	insideSpace Space
	parentSpace *Space
}

func (be *Bag) IsEmpty() bool {
	return (be.insideSpace == nil) || be.Count() == 0
}

func (be *Bag) addState(ns State) {
	be.states = addStateHelper(be.states, ns)
}

func (be *Bag) listStates() []State {
	return be.states
}

func (be *Bag) removeState(str string) {
	be.states = removeStateHelper(be.states, str)
}

//func (be *Bag) MakeContainer() {
//	if be.insideSpace == nil {
//		be.insideSpace = &Content{}
//	}
//}

func (be *Bag) Put(e Entity) {
	if be.insideSpace == nil {
		panic("Not a container!")
	}
	be.insideSpace.Put(e)
}

func (be *Bag) Count() uint {
	if be.insideSpace == nil {
		panic("Not a container!")
	}
	return be.insideSpace.Count()
}

func (be *Bag) SetParent(p *Space) {
	be.parentSpace = p
}

type NPC struct {
	states []State
	parentSpace *Space
}

func (be *NPC) addState(ns State) {
	be.states = addStateHelper(be.states, ns)
}

func (be *NPC) listStates() []State {
	return be.states
}

func (be *NPC) removeState(str string) {
	be.states = removeStateHelper(be.states, str)
}

func (be *NPC) SetParent(p *Space) {
	be.parentSpace = p
}

type Buff struct {
	Name string
	Value string
}

func (b *Buff) Index() string {
	return b.Name
}

func (b *Buff) Dub(str string) {
	b.Name = str
}

func (b *Buff) Get() string {
	return b.Value
}

func (b *Buff) Set(str string) {
	b.Value = str
}

func listStates(t *testing.T, fred *NPC) {
	fred_states := fred.listStates()
	if len(fred_states) == 0 {
		t.Logf("Fred has no states")
		return
	}
	t.Logf("Fred has these states:")
	for _, s := range fred_states {
		t.Logf("%s = %s", s.Index(), s.Get())
	}
}

func TestStateInterface(t *testing.T) {
	// verify (?) that Buff implements State
	var _ State = (*Buff)(nil)
	// etc
	var _ Entity = (*NPC)(nil)
	// etc
	var _ Entity = (*Bag)(nil)
	var _ Space = (*Content)(nil)
	var _ Entity = (*Marble)(nil)
}

// States may interact with Entities and Spaces
// Not with Actions, Forms or other States

func TestStateEntityRelationship(t *testing.T) {
	fred := &NPC{}
	{
		listStates(t, fred)
		mighty := &Buff{"Mighty", "10"}
		fred.addState(mighty)
		listStates(t, fred)
		mighty.Set("40")
		listStates(t, fred)
		fred.addState(&Buff{"Boosted", "+1"})
		fred.addState(&Buff{"Job", "Bee Keeper"})
		fred.addState(&Buff{"Hate", "wasps"})
		fred.addState(&Buff{"Love", "bees"})
	}
	{
		listStates(t, fred)
		fred.removeState("Mighty")
		listStates(t, fred)
	}
}

func TestStateSpaceRelationship(t *testing.T) {
	// A Space can have States
	bag := &Bag{}
	if bag.IsEmpty() {
		t.Logf("Bag is empty")
	}
	bag.insideSpace = &Content{}
	//bag.MakeContainer(Content)
	if bag.IsEmpty() {
		t.Logf("Bag is still empty")
	}
	bag.Put(&Marble{0.5, "catseye", nil, nil})
	bag.Put(&Marble{1.5, "pearly", nil, nil})
	bag.Put(&Marble{0.5, "iridescent", nil, nil})
	t.Logf("The bag has %d marbles.", bag.Count())
}

