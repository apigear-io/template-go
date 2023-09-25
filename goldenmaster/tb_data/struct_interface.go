package data

import (
	"goldenmaster/tb_data/api"
)

type structInterfaceImpl struct {
	api.INotifier
	propBool   api.StructBool
	propInt    api.StructInt
	propFloat  api.StructFloat
	propString api.StructString
}

var _ api.StructInterface = (*structInterfaceImpl)(nil)
var _ api.INotifier = (*structInterfaceImpl)(nil)

func NewStructInterface(notifier api.INotifier) api.StructInterface {
	obj := &structInterfaceImpl{
		INotifier:  notifier,
		propBool:   api.StructBool{},
		propInt:    api.StructInt{},
		propFloat:  api.StructFloat{},
		propString: api.StructString{},
	}
	return obj
}

// property get propBool
func (s *structInterfaceImpl) GetPropBool() api.StructBool {
	return api.StructBool{}
}

// property set propBool
func (s *structInterfaceImpl) SetPropBool(propBool api.StructBool) {
}

// property get propInt
func (s *structInterfaceImpl) GetPropInt() api.StructInt {
	return api.StructInt{}
}

// property set propInt
func (s *structInterfaceImpl) SetPropInt(propInt api.StructInt) {
}

// property get propFloat
func (s *structInterfaceImpl) GetPropFloat() api.StructFloat {
	return api.StructFloat{}
}

// property set propFloat
func (s *structInterfaceImpl) SetPropFloat(propFloat api.StructFloat) {
}

// property get propString
func (s *structInterfaceImpl) GetPropString() api.StructString {
	return api.StructString{}
}

// property set propString
func (s *structInterfaceImpl) SetPropString(propString api.StructString) {
}

// method funcBool
func (s *structInterfaceImpl) FuncBool(paramBool api.StructBool) api.StructBool {
	return api.StructBool{}
}

// method funcInt
func (s *structInterfaceImpl) FuncInt(paramInt api.StructInt) api.StructInt {
	return api.StructInt{}
}

// method funcFloat
func (s *structInterfaceImpl) FuncFloat(paramFloat api.StructFloat) api.StructFloat {
	return api.StructFloat{}
}

// method funcString
func (s *structInterfaceImpl) FuncString(paramString api.StructString) api.StructString {
	return api.StructString{}
}
