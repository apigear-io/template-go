package testbed

import (
	"goldenmaster/tb_adv/api"
)

type nestedStruct2InterfaceImpl struct {
    api.INotifier
    prop1 api.NestedStruct1
    prop2 api.NestedStruct2
}

var _ api.NestedStruct2Interface = (*nestedStruct2InterfaceImpl)(nil)
var _ api.INotifier = (*nestedStruct2InterfaceImpl)(nil)

func NewNestedStruct2Interface(notifier api.INotifier) api.NestedStruct2Interface {
	obj := &nestedStruct2InterfaceImpl{
        INotifier: notifier,
        prop1: api.NestedStruct1{},
        prop2: api.NestedStruct2{},
    }
  	return obj
}
// property get prop1
func (n *nestedStruct2InterfaceImpl) GetProp1() api.NestedStruct1 {
	return n.prop1
}

// property set prop1
func (n *nestedStruct2InterfaceImpl) SetProp1(prop1 api.NestedStruct1) {
    n.prop1 = prop1
    n.NotifyPropertyChanged("prop1", prop1)
}

// property get prop2
func (n *nestedStruct2InterfaceImpl) GetProp2() api.NestedStruct2 {
	return n.prop2
}

// property set prop2
func (n *nestedStruct2InterfaceImpl) SetProp2(prop2 api.NestedStruct2) {
    n.prop2 = prop2
    n.NotifyPropertyChanged("prop2", prop2)
}

