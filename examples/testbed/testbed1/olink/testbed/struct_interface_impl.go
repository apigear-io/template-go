package testbed

import (
	"testbed/testbed1/api"
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
	return s.propBool
}

// property set propBool
func (s *structInterfaceImpl) SetPropBool(propBool api.StructBool) {
	s.propBool = propBool
	s.NotifyPropertyChanged("propBool", propBool)
}

// property get propInt
func (s *structInterfaceImpl) GetPropInt() api.StructInt {
	return s.propInt
}

// property set propInt
func (s *structInterfaceImpl) SetPropInt(propInt api.StructInt) {
	s.propInt = propInt
	s.NotifyPropertyChanged("propInt", propInt)
}

// property get propFloat
func (s *structInterfaceImpl) GetPropFloat() api.StructFloat {
	return s.propFloat
}

// property set propFloat
func (s *structInterfaceImpl) SetPropFloat(propFloat api.StructFloat) {
	s.propFloat = propFloat
	s.NotifyPropertyChanged("propFloat", propFloat)
}

// property get propString
func (s *structInterfaceImpl) GetPropString() api.StructString {
	return s.propString
}

// property set propString
func (s *structInterfaceImpl) SetPropString(propString api.StructString) {
	s.propString = propString
	s.NotifyPropertyChanged("propString", propString)
}

// method funcBool
func (s *structInterfaceImpl) FuncBool(paramBool api.StructBool) api.StructBool {
	s.SetPropBool(paramBool)
	s.NotifySignal("sigBool", []any{paramBool})
	return paramBool
}

// method funcInt
func (s *structInterfaceImpl) FuncInt(paramInt api.StructInt) api.StructInt {
	s.SetPropInt(paramInt)
	s.NotifySignal("sigInt", []any{paramInt})
	return paramInt
}

// method funcFloat
func (s *structInterfaceImpl) FuncFloat(paramFloat api.StructFloat) api.StructFloat {
	s.SetPropFloat(paramFloat)
	s.NotifySignal("sigFloat", []any{paramFloat})
	return paramFloat
}

// method funcString
func (s *structInterfaceImpl) FuncString(paramString api.StructString) api.StructString {
	s.SetPropString(paramString)
	s.NotifySignal("sigString", []any{paramString})
	return paramString
}
