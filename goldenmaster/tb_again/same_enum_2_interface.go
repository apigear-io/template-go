package again

import (
	"goldenmaster/tb_again/api"
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
	return api.Enum1Value1
}

// property set prop1
func (s *sameEnum2InterfaceImpl) SetProp1(prop1 api.Enum1) {
}

// property get prop2
func (s *sameEnum2InterfaceImpl) GetProp2() api.Enum2 {
	return api.Enum2Value1
}

// property set prop2
func (s *sameEnum2InterfaceImpl) SetProp2(prop2 api.Enum2) {
}
