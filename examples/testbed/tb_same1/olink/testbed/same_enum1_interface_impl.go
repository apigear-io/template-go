package testbed

import (
	"testbed/tb_same1/api"
)

type sameEnum1InterfaceImpl struct {
	api.INotifier
	prop1 api.Enum1
}

var _ api.SameEnum1Interface = (*sameEnum1InterfaceImpl)(nil)
var _ api.INotifier = (*sameEnum1InterfaceImpl)(nil)

func NewSameEnum1Interface(notifier api.INotifier) api.SameEnum1Interface {
	obj := &sameEnum1InterfaceImpl{
		INotifier: notifier,
		prop1:     api.Enum1Value1,
	}
	return obj
}

// property get prop1
func (s *sameEnum1InterfaceImpl) GetProp1() api.Enum1 {
	return s.prop1
}

// property set prop1
func (s *sameEnum1InterfaceImpl) SetProp1(prop1 api.Enum1) {
	s.prop1 = prop1
	s.NotifyPropertyChanged("prop1", prop1)
}

// method func1
func (s *sameEnum1InterfaceImpl) Func1(param1 api.Enum1) api.Enum1 {
	s.SetProp1(param1)
	s.NotifySignal("sig1", []any{param1})
	return param1
}
