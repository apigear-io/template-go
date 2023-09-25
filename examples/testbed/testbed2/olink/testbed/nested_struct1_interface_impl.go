package testbed

import (
	"testbed/testbed2/api"
)

type nestedStruct1InterfaceImpl struct {
	api.INotifier
	prop1 api.NestedStruct1
}

var _ api.NestedStruct1Interface = (*nestedStruct1InterfaceImpl)(nil)
var _ api.INotifier = (*nestedStruct1InterfaceImpl)(nil)

func NewNestedStruct1Interface(notifier api.INotifier) api.NestedStruct1Interface {
	obj := &nestedStruct1InterfaceImpl{
		INotifier: notifier,
		prop1:     api.NestedStruct1{},
	}
	return obj
}

// property get prop1
func (n *nestedStruct1InterfaceImpl) GetProp1() api.NestedStruct1 {
	return n.prop1
}

// property set prop1
func (n *nestedStruct1InterfaceImpl) SetProp1(prop1 api.NestedStruct1) {
	n.prop1 = prop1
	n.NotifyPropertyChanged("prop1", prop1)
}

// method func1
func (n *nestedStruct1InterfaceImpl) Func1(param1 api.NestedStruct1) api.NestedStruct1 {
	n.SetProp1(param1)
	n.NotifySignal("sig1", []any{param1})
	return param1
}
