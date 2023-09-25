package testbed

import (
	"testbed/tb_simple/api"
)

type simpleArrayInterfaceImpl struct {
	api.INotifier
	propBool   []bool
	propInt    []int32
	propFloat  []float32
	propString []string
}

var _ api.SimpleArrayInterface = (*simpleArrayInterfaceImpl)(nil)
var _ api.INotifier = (*simpleArrayInterfaceImpl)(nil)

func NewSimpleArrayInterface(notifier api.INotifier) api.SimpleArrayInterface {
	obj := &simpleArrayInterfaceImpl{
		INotifier:  notifier,
		propBool:   make([]bool, 0),
		propInt:    make([]int32, 0),
		propFloat:  make([]float32, 0),
		propString: make([]string, 0),
	}
	return obj
}

// property get propBool
func (s *simpleArrayInterfaceImpl) GetPropBool() []bool {
	return s.propBool
}

// property set propBool
func (s *simpleArrayInterfaceImpl) SetPropBool(propBool []bool) {
	s.propBool = propBool
	s.NotifyPropertyChanged("propBool", propBool)
}

// property get propInt
func (s *simpleArrayInterfaceImpl) GetPropInt() []int32 {
	return s.propInt
}

// property set propInt
func (s *simpleArrayInterfaceImpl) SetPropInt(propInt []int32) {
	s.propInt = propInt
	s.NotifyPropertyChanged("propInt", propInt)
}

// property get propFloat
func (s *simpleArrayInterfaceImpl) GetPropFloat() []float32 {
	return s.propFloat
}

// property set propFloat
func (s *simpleArrayInterfaceImpl) SetPropFloat(propFloat []float32) {
	s.propFloat = propFloat
	s.NotifyPropertyChanged("propFloat", propFloat)
}

// property get propString
func (s *simpleArrayInterfaceImpl) GetPropString() []string {
	return s.propString
}

// property set propString
func (s *simpleArrayInterfaceImpl) SetPropString(propString []string) {
	s.propString = propString
	s.NotifyPropertyChanged("propString", propString)
}

// method funcBool
func (s *simpleArrayInterfaceImpl) FuncBool(paramBool []bool) []bool {
	s.SetPropBool(paramBool)
	s.NotifySignal("sigBool", []any{paramBool})
	return paramBool
}

// method funcInt
func (s *simpleArrayInterfaceImpl) FuncInt(paramInt []int32) []int32 {
	s.SetPropInt(paramInt)
	s.NotifySignal("sigInt", []any{paramInt})
	return paramInt
}

// method funcFloat
func (s *simpleArrayInterfaceImpl) FuncFloat(paramFloat []float32) []float32 {
	s.SetPropFloat(paramFloat)
	s.NotifySignal("sigFloat", []any{paramFloat})
	return paramFloat
}

// method funcString
func (s *simpleArrayInterfaceImpl) FuncString(paramString []string) []string {
	s.SetPropString(paramString)
	s.NotifySignal("sigString", []any{paramString})
	return paramString
}
