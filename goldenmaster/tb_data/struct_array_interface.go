package data

import (
	"goldenmaster/tb_data/api"
)

type structArrayInterfaceImpl struct {
	api.INotifier
	propBool   []api.StructBool
	propInt    []api.StructInt
	propFloat  []api.StructFloat
	propString []api.StructString
}

var _ api.StructArrayInterface = (*structArrayInterfaceImpl)(nil)
var _ api.INotifier = (*structArrayInterfaceImpl)(nil)

func NewStructArrayInterface(notifier api.INotifier) api.StructArrayInterface {
	obj := &structArrayInterfaceImpl{
		INotifier:  notifier,
		propBool:   make([]api.StructBool, 0),
		propInt:    make([]api.StructInt, 0),
		propFloat:  make([]api.StructFloat, 0),
		propString: make([]api.StructString, 0),
	}
	return obj
}

// property get propBool
func (s *structArrayInterfaceImpl) GetPropBool() []api.StructBool {
	return []api.StructBool{}
}

// property set propBool
func (s *structArrayInterfaceImpl) SetPropBool(propBool []api.StructBool) {
}

// property get propInt
func (s *structArrayInterfaceImpl) GetPropInt() []api.StructInt {
	return []api.StructInt{}
}

// property set propInt
func (s *structArrayInterfaceImpl) SetPropInt(propInt []api.StructInt) {
}

// property get propFloat
func (s *structArrayInterfaceImpl) GetPropFloat() []api.StructFloat {
	return []api.StructFloat{}
}

// property set propFloat
func (s *structArrayInterfaceImpl) SetPropFloat(propFloat []api.StructFloat) {
}

// property get propString
func (s *structArrayInterfaceImpl) GetPropString() []api.StructString {
	return []api.StructString{}
}

// property set propString
func (s *structArrayInterfaceImpl) SetPropString(propString []api.StructString) {
}

// method funcBool
func (s *structArrayInterfaceImpl) FuncBool(paramBool []api.StructBool) []api.StructBool {
	return []api.StructBool{}
}

// method funcInt
func (s *structArrayInterfaceImpl) FuncInt(paramInt []api.StructInt) []api.StructInt {
	return []api.StructInt{}
}

// method funcFloat
func (s *structArrayInterfaceImpl) FuncFloat(paramFloat []api.StructFloat) []api.StructFloat {
	return []api.StructFloat{}
}

// method funcString
func (s *structArrayInterfaceImpl) FuncString(paramString []api.StructString) []api.StructString {
	return []api.StructString{}
}
