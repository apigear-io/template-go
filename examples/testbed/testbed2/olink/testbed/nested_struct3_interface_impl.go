package testbed

import (
	"testbed/testbed2/api"
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
	return n.prop1
}

// property set prop1
func (n *nestedStruct3InterfaceImpl) SetProp1(prop1 api.NestedStruct1) {
	n.prop1 = prop1
	n.NotifyPropertyChanged("prop1", prop1)
}

// property get prop2
func (n *nestedStruct3InterfaceImpl) GetProp2() api.NestedStruct2 {
	return n.prop2
}

// property set prop2
func (n *nestedStruct3InterfaceImpl) SetProp2(prop2 api.NestedStruct2) {
	n.prop2 = prop2
	n.NotifyPropertyChanged("prop2", prop2)
}

// property get prop3
func (n *nestedStruct3InterfaceImpl) GetProp3() api.NestedStruct3 {
	return n.prop3
}

// property set prop3
func (n *nestedStruct3InterfaceImpl) SetProp3(prop3 api.NestedStruct3) {
	n.prop3 = prop3
	n.NotifyPropertyChanged("prop3", prop3)
}

// method func1
func (n *nestedStruct3InterfaceImpl) Func1(param1 api.NestedStruct1) api.NestedStruct1 {
	n.SetProp1(param1)
	n.NotifySignal("sig1", []any{param1})
	return param1
}

// method func2
func (n *nestedStruct3InterfaceImpl) Func2(param1 api.NestedStruct1, param2 api.NestedStruct2) api.NestedStruct1 {
	n.SetProp1(param1)
	n.NotifySignal("sig1", []any{param1})
	n.SetProp2(param2)
	n.NotifySignal("sig2", []any{param2})
	return param1
}

// method func3
func (n *nestedStruct3InterfaceImpl) Func3(param1 api.NestedStruct1, param2 api.NestedStruct2, param3 api.NestedStruct3) api.NestedStruct1 {
	n.SetProp1(param1)
	n.NotifySignal("sig1", []any{param1})
	n.SetProp2(param2)
	n.NotifySignal("sig2", []any{param2})
	n.SetProp3(param3)
	n.NotifySignal("sig3", []any{param3})
	return param1
}
