package again

import (
	"goldenmaster/tb_again/api"
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
		prop1:     api.Struct2{},
		prop2:     api.Struct2{},
	}
	return obj
}

// property get prop1
func (s *sameStruct2InterfaceImpl) GetProp1() api.Struct2 {
	return api.Struct2{}
}

// property set prop1
func (s *sameStruct2InterfaceImpl) SetProp1(prop1 api.Struct2) {
}

// property get prop2
func (s *sameStruct2InterfaceImpl) GetProp2() api.Struct2 {
	return api.Struct2{}
}

// property set prop2
func (s *sameStruct2InterfaceImpl) SetProp2(prop2 api.Struct2) {
}
