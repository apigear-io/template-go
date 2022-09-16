package testbed

import (
	"goldenmaster/tb_enum/api"
)

type enumInterfaceImpl struct {
    api.INotifier
    prop0 api.Enum0
    prop1 api.Enum1
    prop2 api.Enum2
    prop3 api.Enum3
}

var _ api.EnumInterface = (*enumInterfaceImpl)(nil)
var _ api.INotifier = (*enumInterfaceImpl)(nil)

func NewEnumInterface(notifier api.INotifier) api.EnumInterface {
	obj := &enumInterfaceImpl{
        INotifier: notifier,
        prop0: api.Enum0Value0,
        prop1: api.Enum1Value1,
        prop2: api.Enum2Value2,
        prop3: api.Enum3Value3,
    }
  	return obj
}
// property get prop0
func (e *enumInterfaceImpl) GetProp0() api.Enum0 {
	return e.prop0
}

// property set prop0
func (e *enumInterfaceImpl) SetProp0(prop0 api.Enum0) {
    e.prop0 = prop0
    e.NotifyPropertyChanged("prop0", prop0)
}

// property get prop1
func (e *enumInterfaceImpl) GetProp1() api.Enum1 {
	return e.prop1
}

// property set prop1
func (e *enumInterfaceImpl) SetProp1(prop1 api.Enum1) {
    e.prop1 = prop1
    e.NotifyPropertyChanged("prop1", prop1)
}

// property get prop2
func (e *enumInterfaceImpl) GetProp2() api.Enum2 {
	return e.prop2
}

// property set prop2
func (e *enumInterfaceImpl) SetProp2(prop2 api.Enum2) {
    e.prop2 = prop2
    e.NotifyPropertyChanged("prop2", prop2)
}

// property get prop3
func (e *enumInterfaceImpl) GetProp3() api.Enum3 {
	return e.prop3
}

// property set prop3
func (e *enumInterfaceImpl) SetProp3(prop3 api.Enum3) {
    e.prop3 = prop3
    e.NotifyPropertyChanged("prop3", prop3)
}

// method func0
func (e *enumInterfaceImpl) Func0(param0 api.Enum0) api.Enum0 {
    e.SetProp0(param0)
    e.NotifySignal("sig0", []any{ param0 })
    return param0
}
    

// method func1
func (e *enumInterfaceImpl) Func1(param1 api.Enum1) api.Enum1 {
    e.SetProp1(param1)
    e.NotifySignal("sig1", []any{ param1 })
    return param1
}
    

// method func2
func (e *enumInterfaceImpl) Func2(param2 api.Enum2) api.Enum2 {
    e.SetProp2(param2)
    e.NotifySignal("sig2", []any{ param2 })
    return param2
}
    

// method func3
func (e *enumInterfaceImpl) Func3(param3 api.Enum3) api.Enum3 {
    e.SetProp3(param3)
    e.NotifySignal("sig3", []any{ param3 })
    return param3
}
    

