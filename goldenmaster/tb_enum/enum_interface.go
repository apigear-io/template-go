package enum

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
	return api.Enum0Value0
}

// property set prop0
func (e *enumInterfaceImpl) SetProp0(prop0 api.Enum0) {  
}

// property get prop1
func (e *enumInterfaceImpl) GetProp1() api.Enum1 {
	return api.Enum1Value1
}

// property set prop1
func (e *enumInterfaceImpl) SetProp1(prop1 api.Enum1) {  
}

// property get prop2
func (e *enumInterfaceImpl) GetProp2() api.Enum2 {
	return api.Enum2Value2
}

// property set prop2
func (e *enumInterfaceImpl) SetProp2(prop2 api.Enum2) {  
}

// property get prop3
func (e *enumInterfaceImpl) GetProp3() api.Enum3 {
	return api.Enum3Value3
}

// property set prop3
func (e *enumInterfaceImpl) SetProp3(prop3 api.Enum3) {  
}

// method func0
func (e *enumInterfaceImpl) Func0(param0 api.Enum0) api.Enum0 {
  return api.Enum0Value0
}
    

// method func1
func (e *enumInterfaceImpl) Func1(param1 api.Enum1) api.Enum1 {
  return api.Enum1Value1
}
    

// method func2
func (e *enumInterfaceImpl) Func2(param2 api.Enum2) api.Enum2 {
  return api.Enum2Value2
}
    

// method func3
func (e *enumInterfaceImpl) Func3(param3 api.Enum3) api.Enum3 {
  return api.Enum3Value3
}
    

