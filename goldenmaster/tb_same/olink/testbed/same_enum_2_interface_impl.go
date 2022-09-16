package testbed

import (
	"goldenmaster/tb_same/api"
)

type sameEnum2InterfaceImpl struct {
    api.INotifier
    prop1 api.Enum1
    prop2 api.Enum2
}

var _ api.SameEnum2Interface = (*sameEnum2InterfaceImpl)(nil)
var _ api.INotifier = (*sameEnum2InterfaceImpl)(nil)

func NewSameEnum2Interface(notifier api.INotifier) api.SameEnum2Interface {
	obj := &sameEnum2InterfaceImpl{
        INotifier: notifier,
        prop1: api.Enum1Value1,
        prop2: api.Enum2Value1,
    }
  	return obj
}
// property get prop1
func (s *sameEnum2InterfaceImpl) GetProp1() api.Enum1 {
	return s.prop1
}

// property set prop1
func (s *sameEnum2InterfaceImpl) SetProp1(prop1 api.Enum1) {
    s.prop1 = prop1
    s.NotifyPropertyChanged("prop1", prop1)
}

// property get prop2
func (s *sameEnum2InterfaceImpl) GetProp2() api.Enum2 {
	return s.prop2
}

// property set prop2
func (s *sameEnum2InterfaceImpl) SetProp2(prop2 api.Enum2) {
    s.prop2 = prop2
    s.NotifyPropertyChanged("prop2", prop2)
}

