package testbed

import (
	"goldenmaster/tb_same/api"
)

type sameStruct1InterfaceImpl struct {
	api.INotifier
	prop1 api.Struct1
}

var _ api.SameStruct1Interface = (*sameStruct1InterfaceImpl)(nil)
var _ api.INotifier = (*sameStruct1InterfaceImpl)(nil)

func NewSameStruct1Interface(notifier api.INotifier) api.SameStruct1Interface {
	obj := &sameStruct1InterfaceImpl{
		INotifier: notifier,
		prop1:     api.Struct1{},
	}
	return obj
}

// property get prop1
func (s *sameStruct1InterfaceImpl) GetProp1() api.Struct1 {
	return s.prop1
}

// property set prop1
func (s *sameStruct1InterfaceImpl) SetProp1(prop1 api.Struct1) {
	s.prop1 = prop1
	s.NotifyPropertyChanged("prop1", prop1)
}
