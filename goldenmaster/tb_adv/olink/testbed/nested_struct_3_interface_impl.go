package testbed

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
        prop1: api.NestedStruct1{},
        prop2: api.NestedStruct2{},
        prop3: api.NestedStruct3{},
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

