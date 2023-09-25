package api

// enum Enum0
type Enum0 int

const (
	Enum0_UNKNOWN Enum0 = -1
	Enum0Value0   Enum0 = 0
	Enum0Value1   Enum0 = 1
	Enum0Value2   Enum0 = 2
)

// enum Enum1
type Enum1 int

const (
	Enum1_UNKNOWN Enum1 = -1
	Enum1Value1   Enum1 = 1
	Enum1Value2   Enum1 = 2
	Enum1Value3   Enum1 = 3
)

// enum Enum2
type Enum2 int

const (
	Enum2_UNKNOWN Enum2 = -1
	Enum2Value2   Enum2 = 2
	Enum2Value1   Enum2 = 1
	Enum2Value0   Enum2 = 0
)

// enum Enum3
type Enum3 int

const (
	Enum3_UNKNOWN Enum3 = -1
	Enum3Value3   Enum3 = 3
	Enum3Value2   Enum3 = 2
	Enum3Value1   Enum3 = 1
)

// Data Structures

// Object Interfaces

type INotifier interface {
	NotifyPropertyChanged(name string, value any)
	NotifySignal(name string, args []any)
}
type EnumInterface interface {
	INotifier
	// Properties
	GetProp0() Enum0
	SetProp0(prop0 Enum0)
	GetProp1() Enum1
	SetProp1(prop1 Enum1)
	GetProp2() Enum2
	SetProp2(prop2 Enum2)
	GetProp3() Enum3
	SetProp3(prop3 Enum3)
	// Methods
	Func0(param0 Enum0) Enum0
	Func1(param1 Enum1) Enum1
	Func2(param2 Enum2) Enum2
	Func3(param3 Enum3) Enum3
}
