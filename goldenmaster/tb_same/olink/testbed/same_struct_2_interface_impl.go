package testbed

import (
	"goldenmaster/tb_same/api"
)

type sameStruct2InterfaceImpl struct {
    api.INotifier
    prop1 api.Struct2
    prop2 api.Struct2
}

var _ api.SameStruct2Interface = (*sameStruct2InterfaceImpl)(nil)
var _ api.INotifier = (*sameStruct2InterfaceImpl)(nil)

func NewSameStruct2Interface(notifier api.INotifier) api.SameStruct2Interface {
	obj := &sameStruct2InterfaceImpl{
        INotifier: notifier,
        prop1: api.Struct2{},
        prop2: api.Struct2{},
    }
  	return obj
}
// property get prop1
func (s *sameStruct2InterfaceImpl) GetProp1() api.Struct2 {
	return s.prop1
}

// property set prop1
func (s *sameStruct2InterfaceImpl) SetProp1(prop1 api.Struct2) {
    s.prop1 = prop1
    s.NotifyPropertyChanged("prop1", prop1)
}

// property get prop2
func (s *sameStruct2InterfaceImpl) GetProp2() api.Struct2 {
	return s.prop2
}

// property set prop2
func (s *sameStruct2InterfaceImpl) SetProp2(prop2 api.Struct2) {
    s.prop2 = prop2
    s.NotifyPropertyChanged("prop2", prop2)
}

