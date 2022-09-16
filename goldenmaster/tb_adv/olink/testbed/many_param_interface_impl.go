package testbed

import (
	"goldenmaster/tb_adv/api"
)

type manyParamInterfaceImpl struct {
    api.INotifier
    prop1 int64
    prop2 int64
    prop3 int64
    prop4 int64
}

var _ api.ManyParamInterface = (*manyParamInterfaceImpl)(nil)
var _ api.INotifier = (*manyParamInterfaceImpl)(nil)

func NewManyParamInterface(notifier api.INotifier) api.ManyParamInterface {
	obj := &manyParamInterfaceImpl{
        INotifier: notifier,
        prop1: int64(0),
        prop2: int64(0),
        prop3: int64(0),
        prop4: int64(0),
    }
  	return obj
}
// property get prop1
func (m *manyParamInterfaceImpl) GetProp1() int64 {
	return m.prop1
}

// property set prop1
func (m *manyParamInterfaceImpl) SetProp1(prop1 int64) {
    m.prop1 = prop1
    m.NotifyPropertyChanged("prop1", prop1)
}

// property get prop2
func (m *manyParamInterfaceImpl) GetProp2() int64 {
	return m.prop2
}

// property set prop2
func (m *manyParamInterfaceImpl) SetProp2(prop2 int64) {
    m.prop2 = prop2
    m.NotifyPropertyChanged("prop2", prop2)
}

// property get prop3
func (m *manyParamInterfaceImpl) GetProp3() int64 {
	return m.prop3
}

// property set prop3
func (m *manyParamInterfaceImpl) SetProp3(prop3 int64) {
    m.prop3 = prop3
    m.NotifyPropertyChanged("prop3", prop3)
}

// property get prop4
func (m *manyParamInterfaceImpl) GetProp4() int64 {
	return m.prop4
}

// property set prop4
func (m *manyParamInterfaceImpl) SetProp4(prop4 int64) {
    m.prop4 = prop4
    m.NotifyPropertyChanged("prop4", prop4)
}

