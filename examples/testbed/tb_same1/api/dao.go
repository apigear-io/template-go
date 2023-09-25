package api

// func SameStruct1Interface.func1(param1 Struct1) Struct1

type SameStruct1InterfaceFunc1Request struct {
	Param1 Struct1 `json:"param1"`
}

type SameStruct1InterfaceFunc1Reply struct {
}

// func SameStruct2Interface.func1(param1 Struct1) Struct1

type SameStruct2InterfaceFunc1Request struct {
	Param1 Struct1 `json:"param1"`
}

type SameStruct2InterfaceFunc1Reply struct {
}

// func SameStruct2Interface.func2(param1 Struct1, param2 Struct2) Struct1

type SameStruct2InterfaceFunc2Request struct {
	Param1 Struct1 `json:"param1"`
	Param2 Struct2 `json:"param2"`
}

type SameStruct2InterfaceFunc2Reply struct {
}

// func SameEnum1Interface.func1(param1 Enum1) Enum1

type SameEnum1InterfaceFunc1Request struct {
	Param1 Enum1 `json:"param1"`
}

type SameEnum1InterfaceFunc1Reply struct {
}

// func SameEnum2Interface.func1(param1 Enum1) Enum1

type SameEnum2InterfaceFunc1Request struct {
	Param1 Enum1 `json:"param1"`
}

type SameEnum2InterfaceFunc1Reply struct {
}

// func SameEnum2Interface.func2(param1 Enum1, param2 Enum2) Enum1

type SameEnum2InterfaceFunc2Request struct {
	Param1 Enum1 `json:"param1"`
	Param2 Enum2 `json:"param2"`
}

type SameEnum2InterfaceFunc2Reply struct {
}
