package adv

import (
	"goldenmaster/tb_adv/api"
)

type nestedStruct3InterfaceImpl struct {
	api.INotifier
	prop1 api.NestedStruct1
	prop2 api.NestedStruct2
	prop3 api.NestedStruct3
}

var _ api.NestedStruct3Interface = (*nestedStruct3InterfaceImpl)(nil)
var _ api.INotifier = (*nestedStruct3InterfaceImpl)(nil)

func NewNestedStruct3Interface(notifier api.INotifier) api.NestedStruct3Interface {
	obj := &nestedStruct3InterfaceImpl{
		INotifier: notifier,
		prop1:     api.NestedStruct1{},
		prop2:     api.NestedStruct2{},
		prop3:     api.NestedStruct3{},
	}
	return obj
}

// property get prop1
func (n *nestedStruct3InterfaceImpl) GetProp1() api.NestedStruct1 {
	return api.NestedStruct1{}
}

// property set prop1
func (n *nestedStruct3InterfaceImpl) SetProp1(prop1 api.NestedStruct1) {
}

// property get prop2
func (n *nestedStruct3InterfaceImpl) GetProp2() api.NestedStruct2 {
	return api.NestedStruct2{}
}

// property set prop2
func (n *nestedStruct3InterfaceImpl) SetProp2(prop2 api.NestedStruct2) {
}

// property get prop3
func (n *nestedStruct3InterfaceImpl) GetProp3() api.NestedStruct3 {
	return api.NestedStruct3{}
}

// property set prop3
func (n *nestedStruct3InterfaceImpl) SetProp3(prop3 api.NestedStruct3) {
}
