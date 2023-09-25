package testbed

import (
	"testbed/tb_same1/api"
)

type sameStruct2InterfaceImpl struct {
	api.INotifier
	prop1 api.Struct1
	prop2 api.Struct2
}

var _ api.SameStruct2Interface = (*sameStruct2InterfaceImpl)(nil)
var _ api.INotifier = (*sameStruct2InterfaceImpl)(nil)

func NewSameStruct2Interface(notifier api.INotifier) api.SameStruct2Interface {
	obj := &sameStruct2InterfaceImpl{
		INotifier: notifier,
		prop1:     api.Struct1{},
		prop2:     api.Struct2{},
	}
	return obj
}

// property get prop1
func (s *sameStruct2InterfaceImpl) GetProp1() api.Struct1 {
	return s.prop1
}

// property set prop1
func (s *sameStruct2InterfaceImpl) SetProp1(prop1 api.Struct1) {
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

// method func1
func (s *sameStruct2InterfaceImpl) Func1(param1 api.Struct1) api.Struct1 {
	s.SetProp1(param1)
	s.NotifySignal("sig1", []any{param1})
	return param1
}

// method func2
func (s *sameStruct2InterfaceImpl) Func2(param1 api.Struct1, param2 api.Struct2) api.Struct1 {
	s.SetProp1(param1)
	s.NotifySignal("sig1", []any{param1})
	s.SetProp2(param2)
	s.NotifySignal("sig2", []any{param2})
	return param1
}
