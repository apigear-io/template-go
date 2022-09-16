package simple

import (
	"goldenmaster/tb_simple/api"
)

type simpleInterfaceImpl struct {
    api.INotifier
    propBool bool
    propInt int64
    propFloat float64
    propString string
}

var _ api.SimpleInterface = (*simpleInterfaceImpl)(nil)
var _ api.INotifier = (*simpleInterfaceImpl)(nil)

func NewSimpleInterface(notifier api.INotifier) api.SimpleInterface {
	obj := &simpleInterfaceImpl{
        INotifier: notifier,
        propBool: false,
        propInt: int64(0),
        propFloat: float64(0.0),
        propString: "",
    }
  	return obj
}
// property get propBool
func (s *simpleInterfaceImpl) GetPropBool() bool {
	return false
}

// property set propBool
func (s *simpleInterfaceImpl) SetPropBool(propBool bool) {  
}

// property get propInt
func (s *simpleInterfaceImpl) GetPropInt() int64 {
	return int64(0)
}

// property set propInt
func (s *simpleInterfaceImpl) SetPropInt(propInt int64) {  
}

// property get propFloat
func (s *simpleInterfaceImpl) GetPropFloat() float64 {
	return float64(0.0)
}

// property set propFloat
func (s *simpleInterfaceImpl) SetPropFloat(propFloat float64) {  
}

// property get propString
func (s *simpleInterfaceImpl) GetPropString() string {
	return ""
}

// property set propString
func (s *simpleInterfaceImpl) SetPropString(propString string) {  
}

// method funcBool
func (s *simpleInterfaceImpl) FuncBool(paramBool bool) bool {
  return false
}
    

// method funcInt
func (s *simpleInterfaceImpl) FuncInt(paramInt int64) int64 {
  return int64(0)
}
    

// method funcFloat
func (s *simpleInterfaceImpl) FuncFloat(paramFloat float64) float64 {
  return float64(0.0)
}
    

// method funcString
func (s *simpleInterfaceImpl) FuncString(paramString string) string {
  return ""
}
    

