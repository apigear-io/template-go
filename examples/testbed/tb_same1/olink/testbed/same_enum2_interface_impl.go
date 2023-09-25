package testbed

import (
	"testbed/tb_same1/api"
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
		prop1:     api.Enum1Value1,
		prop2:     api.Enum2Value1,
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

// method func1
func (s *sameEnum2InterfaceImpl) Func1(param1 api.Enum1) api.Enum1 {
	s.SetProp1(param1)
	s.NotifySignal("sig1", []any{param1})
	return param1
}

// method func2
func (s *sameEnum2InterfaceImpl) Func2(param1 api.Enum1, param2 api.Enum2) api.Enum1 {
	s.SetProp1(param1)
	s.NotifySignal("sig1", []any{param1})
	s.SetProp2(param2)
	s.NotifySignal("sig2", []any{param2})
	return param1
}
