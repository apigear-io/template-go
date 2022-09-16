package testbed

import (
	"goldenmaster/tb_data/api"
)

type structArrayInterfaceImpl struct {
    api.INotifier
    propBool []api.StructBool
    propInt []api.StructInt
    propFloat []api.StructFloat
    propString []api.StructString
}

var _ api.StructArrayInterface = (*structArrayInterfaceImpl)(nil)
var _ api.INotifier = (*structArrayInterfaceImpl)(nil)

func NewStructArrayInterface(notifier api.INotifier) api.StructArrayInterface {
	obj := &structArrayInterfaceImpl{
        INotifier: notifier,
        propBool: make([]api.StructBool, 0),
        propInt: make([]api.StructInt, 0),
        propFloat: make([]api.StructFloat, 0),
        propString: make([]api.StructString, 0),
    }
  	return obj
}
// property get propBool
func (s *structArrayInterfaceImpl) GetPropBool() []api.StructBool {
	return s.propBool
}

// property set propBool
func (s *structArrayInterfaceImpl) SetPropBool(propBool []api.StructBool) {
    s.propBool = propBool
    s.NotifyPropertyChanged("propBool", propBool)
}

// property get propInt
func (s *structArrayInterfaceImpl) GetPropInt() []api.StructInt {
	return s.propInt
}

// property set propInt
func (s *structArrayInterfaceImpl) SetPropInt(propInt []api.StructInt) {
    s.propInt = propInt
    s.NotifyPropertyChanged("propInt", propInt)
}

// property get propFloat
func (s *structArrayInterfaceImpl) GetPropFloat() []api.StructFloat {
	return s.propFloat
}

// property set propFloat
func (s *structArrayInterfaceImpl) SetPropFloat(propFloat []api.StructFloat) {
    s.propFloat = propFloat
    s.NotifyPropertyChanged("propFloat", propFloat)
}

// property get propString
func (s *structArrayInterfaceImpl) GetPropString() []api.StructString {
	return s.propString
}

// property set propString
func (s *structArrayInterfaceImpl) SetPropString(propString []api.StructString) {
    s.propString = propString
    s.NotifyPropertyChanged("propString", propString)
}

// method funcBool
func (s *structArrayInterfaceImpl) FuncBool(paramBool []api.StructBool) []api.StructBool {
    s.SetPropBool(paramBool)
    s.NotifySignal("sigBool", []any{ paramBool })
    return paramBool
}
    

// method funcInt
func (s *structArrayInterfaceImpl) FuncInt(paramInt []api.StructInt) []api.StructInt {
    s.SetPropInt(paramInt)
    s.NotifySignal("sigInt", []any{ paramInt })
    return paramInt
}
    

// method funcFloat
func (s *structArrayInterfaceImpl) FuncFloat(paramFloat []api.StructFloat) []api.StructFloat {
    s.SetPropFloat(paramFloat)
    s.NotifySignal("sigFloat", []any{ paramFloat })
    return paramFloat
}
    

// method funcString
func (s *structArrayInterfaceImpl) FuncString(paramString []api.StructString) []api.StructString {
    s.SetPropString(paramString)
    s.NotifySignal("sigString", []any{ paramString })
    return paramString
}
    

