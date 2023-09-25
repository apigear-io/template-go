package testbed

import (
	"testbed/tb_simple/api"
)

type simpleInterfaceImpl struct {
	api.INotifier
	propBool   bool
	propInt    int32
	propFloat  float32
	propString string
}

var _ api.SimpleInterface = (*simpleInterfaceImpl)(nil)
var _ api.INotifier = (*simpleInterfaceImpl)(nil)

func NewSimpleInterface(notifier api.INotifier) api.SimpleInterface {
	obj := &simpleInterfaceImpl{
		INotifier:  notifier,
		propBool:   false,
		propInt:    int32(0),
		propFloat:  float32(0.0),
		propString: "",
	}
	return obj
}

// property get propBool
func (s *simpleInterfaceImpl) GetPropBool() bool {
	return s.propBool
}

// property set propBool
func (s *simpleInterfaceImpl) SetPropBool(propBool bool) {
	s.propBool = propBool
	s.NotifyPropertyChanged("propBool", propBool)
}

// property get propInt
func (s *simpleInterfaceImpl) GetPropInt() int32 {
	return s.propInt
}

// property set propInt
func (s *simpleInterfaceImpl) SetPropInt(propInt int32) {
	s.propInt = propInt
	s.NotifyPropertyChanged("propInt", propInt)
}

// property get propFloat
func (s *simpleInterfaceImpl) GetPropFloat() float32 {
	return s.propFloat
}

// property set propFloat
func (s *simpleInterfaceImpl) SetPropFloat(propFloat float32) {
	s.propFloat = propFloat
	s.NotifyPropertyChanged("propFloat", propFloat)
}

// property get propString
func (s *simpleInterfaceImpl) GetPropString() string {
	return s.propString
}

// property set propString
func (s *simpleInterfaceImpl) SetPropString(propString string) {
	s.propString = propString
	s.NotifyPropertyChanged("propString", propString)
}

// method funcBool
func (s *simpleInterfaceImpl) FuncBool(paramBool bool) bool {
	s.SetPropBool(paramBool)
	s.NotifySignal("sigBool", []any{paramBool})
	return paramBool
}

// method funcInt
func (s *simpleInterfaceImpl) FuncInt(paramInt int32) int32 {
	s.SetPropInt(paramInt)
	s.NotifySignal("sigInt", []any{paramInt})
	return paramInt
}

// method funcFloat
func (s *simpleInterfaceImpl) FuncFloat(paramFloat float32) float32 {
	s.SetPropFloat(paramFloat)
	s.NotifySignal("sigFloat", []any{paramFloat})
	return paramFloat
}

// method funcString
func (s *simpleInterfaceImpl) FuncString(paramString string) string {
	s.SetPropString(paramString)
	s.NotifySignal("sigString", []any{paramString})
	return paramString
}
