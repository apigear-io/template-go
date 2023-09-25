package api

// enum Enum1
type Enum1 int

const (
	Enum1_UNKNOWN Enum1 = -1
	Enum1Value1   Enum1 = 1
	Enum1Value2   Enum1 = 2
	Enum1Value3   Enum1 = 3
	Enum1Value4   Enum1 = 4
)

// enum Enum2
type Enum2 int

const (
	Enum2_UNKNOWN Enum2 = -1
	Enum2Value1   Enum2 = 1
	Enum2Value2   Enum2 = 2
	Enum2Value3   Enum2 = 3
	Enum2Value4   Enum2 = 4
)

// enum Enum3
type Enum3 int

const (
	Enum3_UNKNOWN Enum3 = -1
	Enum3Value1   Enum3 = 1
	Enum3Value2   Enum3 = 2
	Enum3Value3   Enum3 = 3
	Enum3Value4   Enum3 = 4
)

// Data Structures
type Struct1 struct {
	Field1 int64 `json:"field1"`
}
type Struct2 struct {
	Field1 int64 `json:"field1"`
	Field2 int64 `json:"field2"`
}
type Struct3 struct {
	Field1 int64 `json:"field1"`
	Field2 int64 `json:"field2"`
	Field3 int64 `json:"field3"`
}
type Struct4 struct {
	Field1 int64 `json:"field1"`
	Field2 int64 `json:"field2"`
	Field3 int64 `json:"field3"`
	Field4 int64 `json:"field4"`
}
type NestedStruct1 struct {
	Field1 Struct1 `json:"field1"`
}
type NestedStruct2 struct {
	Field1 Struct1 `json:"field1"`
	Field2 Struct2 `json:"field2"`
}
type NestedStruct3 struct {
	Field1 Struct1 `json:"field1"`
	Field2 Struct2 `json:"field2"`
	Field3 Struct3 `json:"field3"`
}

// Object Interfaces

type INotifier interface {
	NotifyPropertyChanged(name string, value any)
	NotifySignal(name string, args []any)
}
type ManyParamInterface interface {
	INotifier
	// Properties
	GetProp1() int64
	SetProp1(prop1 int64)
	GetProp2() int64
	SetProp2(prop2 int64)
	GetProp3() int64
	SetProp3(prop3 int64)
	GetProp4() int64
	SetProp4(prop4 int64)
	// Methods
}
type NestedStruct1Interface interface {
	INotifier
	// Properties
	GetProp1() NestedStruct1
	SetProp1(prop1 NestedStruct1)
	// Methods
}
type NestedStruct2Interface interface {
	INotifier
	// Properties
	GetProp1() NestedStruct1
	SetProp1(prop1 NestedStruct1)
	GetProp2() NestedStruct2
	SetProp2(prop2 NestedStruct2)
	// Methods
}
type NestedStruct3Interface interface {
	INotifier
	// Properties
	GetProp1() NestedStruct1
	SetProp1(prop1 NestedStruct1)
	GetProp2() NestedStruct2
	SetProp2(prop2 NestedStruct2)
	GetProp3() NestedStruct3
	SetProp3(prop3 NestedStruct3)
	// Methods
}
