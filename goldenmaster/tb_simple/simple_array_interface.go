package simple

import (
	"goldenmaster/tb_simple/api"
)

type simpleArrayInterfaceImpl struct {
	api.INotifier
	propBool   []bool
	propInt    []int64
	propFloat  []float64
	propString []string
}

var _ api.SimpleArrayInterface = (*simpleArrayInterfaceImpl)(nil)
var _ api.INotifier = (*simpleArrayInterfaceImpl)(nil)

func NewSimpleArrayInterface(notifier api.INotifier) api.SimpleArrayInterface {
	obj := &simpleArrayInterfaceImpl{
		INotifier:  notifier,
		propBool:   make([]bool, 0),
		propInt:    make([]int64, 0),
		propFloat:  make([]float64, 0),
		propString: make([]string, 0),
	}
	return obj
}

// property get propBool
func (s *simpleArrayInterfaceImpl) GetPropBool() []bool {
	return []bool{}
}

// property set propBool
func (s *simpleArrayInterfaceImpl) SetPropBool(propBool []bool) {
}

// property get propInt
func (s *simpleArrayInterfaceImpl) GetPropInt() []int64 {
	return []int64{}
}

// property set propInt
func (s *simpleArrayInterfaceImpl) SetPropInt(propInt []int64) {
}

// property get propFloat
func (s *simpleArrayInterfaceImpl) GetPropFloat() []float64 {
	return []float64{}
}

// property set propFloat
func (s *simpleArrayInterfaceImpl) SetPropFloat(propFloat []float64) {
}

// property get propString
func (s *simpleArrayInterfaceImpl) GetPropString() []string {
	return []string{}
}

// property set propString
func (s *simpleArrayInterfaceImpl) SetPropString(propString []string) {
}

// method funcBool
func (s *simpleArrayInterfaceImpl) FuncBool(paramBool []bool) []bool {
	return []bool{}
}

// method funcInt
func (s *simpleArrayInterfaceImpl) FuncInt(paramInt []int64) []int64 {
	return []int64{}
}

// method funcFloat
func (s *simpleArrayInterfaceImpl) FuncFloat(paramFloat []float64) []float64 {
	return []float64{}
}

// method funcString
func (s *simpleArrayInterfaceImpl) FuncString(paramString []string) []string {
	return []string{}
}
