package testbed

import (
	"testbed/testbed2/api"
)

type manyParamInterfaceImpl struct {
	api.INotifier
	prop1 int32
	prop2 int32
	prop3 int32
	prop4 int32
}

var _ api.ManyParamInterface = (*manyParamInterfaceImpl)(nil)
var _ api.INotifier = (*manyParamInterfaceImpl)(nil)

func NewManyParamInterface(notifier api.INotifier) api.ManyParamInterface {
	obj := &manyParamInterfaceImpl{
		INotifier: notifier,
		prop1:     int32(0),
		prop2:     int32(0),
		prop3:     int32(0),
		prop4:     int32(0),
	}
	return obj
}

// property get prop1
func (m *manyParamInterfaceImpl) GetProp1() int32 {
	return m.prop1
}

// property set prop1
func (m *manyParamInterfaceImpl) SetProp1(prop1 int32) {
	m.prop1 = prop1
	m.NotifyPropertyChanged("prop1", prop1)
}

// property get prop2
func (m *manyParamInterfaceImpl) GetProp2() int32 {
	return m.prop2
}

// property set prop2
func (m *manyParamInterfaceImpl) SetProp2(prop2 int32) {
	m.prop2 = prop2
	m.NotifyPropertyChanged("prop2", prop2)
}

// property get prop3
func (m *manyParamInterfaceImpl) GetProp3() int32 {
	return m.prop3
}

// property set prop3
func (m *manyParamInterfaceImpl) SetProp3(prop3 int32) {
	m.prop3 = prop3
	m.NotifyPropertyChanged("prop3", prop3)
}

// property get prop4
func (m *manyParamInterfaceImpl) GetProp4() int32 {
	return m.prop4
}

// property set prop4
func (m *manyParamInterfaceImpl) SetProp4(prop4 int32) {
	m.prop4 = prop4
	m.NotifyPropertyChanged("prop4", prop4)
}

// method func1
func (m *manyParamInterfaceImpl) Func1(param1 int32) int32 {
	m.SetProp1(param1)
	m.NotifySignal("sig1", []any{param1})
	return param1
}

// method func2
func (m *manyParamInterfaceImpl) Func2(param1 int32, param2 int32) int32 {
	m.SetProp1(param1)
	m.NotifySignal("sig1", []any{param1})
	m.SetProp2(param2)
	m.NotifySignal("sig2", []any{param2})
	return param1
}

// method func3
func (m *manyParamInterfaceImpl) Func3(param1 int32, param2 int32, param3 int32) int32 {
	m.SetProp1(param1)
	m.NotifySignal("sig1", []any{param1})
	m.SetProp2(param2)
	m.NotifySignal("sig2", []any{param2})
	m.SetProp3(param3)
	m.NotifySignal("sig3", []any{param3})
	return param1
}

// method func4
func (m *manyParamInterfaceImpl) Func4(param1 int32, param2 int32, param3 int32, param4 int32) int32 {
	m.SetProp1(param1)
	m.NotifySignal("sig1", []any{param1})
	m.SetProp2(param2)
	m.NotifySignal("sig2", []any{param2})
	m.SetProp3(param3)
	m.NotifySignal("sig3", []any{param3})
	m.SetProp4(param4)
	m.NotifySignal("sig4", []any{param4})
	return param1
}
