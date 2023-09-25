package olink

import (
	"testbed/tb_same1/api"
)

type SameStruct1InterfaceProperties struct {
	Prop1 *api.Struct1 `json:"prop1,omitempty"`
}

// func SameStruct1Interface.func1(param1 api.Struct1) api.Struct1

type SameStruct1InterfaceFunc1Request struct {
	Param1 api.Struct1 `json:"param1"`
}

type SameStruct1InterfaceFunc1Reply struct {
}

type SameStruct2InterfaceProperties struct {
	Prop1 *api.Struct1 `json:"prop1,omitempty"`
	Prop2 *api.Struct2 `json:"prop2,omitempty"`
}

// func SameStruct2Interface.func1(param1 api.Struct1) api.Struct1

type SameStruct2InterfaceFunc1Request struct {
	Param1 api.Struct1 `json:"param1"`
}

type SameStruct2InterfaceFunc1Reply struct {
}

// func SameStruct2Interface.func2(param1 api.Struct1, param2 api.Struct2) api.Struct1

type SameStruct2InterfaceFunc2Request struct {
	Param1 api.Struct1 `json:"param1"`
	Param2 api.Struct2 `json:"param2"`
}

type SameStruct2InterfaceFunc2Reply struct {
}

type SameEnum1InterfaceProperties struct {
	Prop1 *api.Enum1 `json:"prop1,omitempty"`
}

// func SameEnum1Interface.func1(param1 api.Enum1) api.Enum1

type SameEnum1InterfaceFunc1Request struct {
	Param1 api.Enum1 `json:"param1"`
}

type SameEnum1InterfaceFunc1Reply struct {
}

type SameEnum2InterfaceProperties struct {
	Prop1 *api.Enum1 `json:"prop1,omitempty"`
	Prop2 *api.Enum2 `json:"prop2,omitempty"`
}

// func SameEnum2Interface.func1(param1 api.Enum1) api.Enum1

type SameEnum2InterfaceFunc1Request struct {
	Param1 api.Enum1 `json:"param1"`
}

type SameEnum2InterfaceFunc1Reply struct {
}

// func SameEnum2Interface.func2(param1 api.Enum1, param2 api.Enum2) api.Enum1

type SameEnum2InterfaceFunc2Request struct {
	Param1 api.Enum1 `json:"param1"`
	Param2 api.Enum2 `json:"param2"`
}

type SameEnum2InterfaceFunc2Reply struct {
}
