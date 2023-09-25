package api

// enum Enum1
type Enum1 int

const (
	Enum1_UNKNOWN Enum1 = -1
	Enum1Value1   Enum1 = 1
	Enum1Value2   Enum1 = 2
)

// enum Enum2
type Enum2 int

const (
	Enum2_UNKNOWN Enum2 = -1
	Enum2Value1   Enum2 = 1
	Enum2Value2   Enum2 = 2
)

// Data Structures
type Struct1 struct {
	Field1 int32 `json:"field1"`
	Field2 int32 `json:"field2"`
	Field3 int32 `json:"field3"`
}
type Struct2 struct {
	Field1 int32 `json:"field1"`
	Field2 int32 `json:"field2"`
	Field3 int32 `json:"field3"`
}

// Object Interfaces

type INotifier interface {
	NotifyPropertyChanged(name string, value any)
	NotifySignal(name string, args []any)
}
type SameStruct1Interface interface {
	INotifier
	// Properties
	GetProp1() Struct1
	SetProp1(prop1 Struct1)
	// Methods
	Func1(param1 Struct1) Struct1
}
type SameStruct2Interface interface {
	INotifier
	// Properties
	GetProp1() Struct1
	SetProp1(prop1 Struct1)
	GetProp2() Struct2
	SetProp2(prop2 Struct2)
	// Methods
	Func1(param1 Struct1) Struct1
	Func2(param1 Struct1, param2 Struct2) Struct1
}
type SameEnum1Interface interface {
	INotifier
	// Properties
	GetProp1() Enum1
	SetProp1(prop1 Enum1)
	// Methods
	Func1(param1 Enum1) Enum1
}
type SameEnum2Interface interface {
	INotifier
	// Properties
	GetProp1() Enum1
	SetProp1(prop1 Enum1)
	GetProp2() Enum2
	SetProp2(prop2 Enum2)
	// Methods
	Func1(param1 Enum1) Enum1
	Func2(param1 Enum1, param2 Enum2) Enum1
}
