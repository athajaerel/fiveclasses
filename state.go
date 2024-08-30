package tbd

import ()

type State interface {
	Index() string
	Dub(string)
	Get() string
	Set(string)
	// later
	//GetInt() int
	//SetInt(int)
	//GetFloat() float32
	//SetFloat(float32)
}

