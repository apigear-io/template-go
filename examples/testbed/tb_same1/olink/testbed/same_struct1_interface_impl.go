package testbed

import (
	"testbed/tb_same1/api"
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

// method func1
func (s *sameStruct1InterfaceImpl) Func1(param1 api.Struct1) api.Struct1 {
	s.SetProp1(param1)
	s.NotifySignal("sig1", []any{param1})
	return param1
}
