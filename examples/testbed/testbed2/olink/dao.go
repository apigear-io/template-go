package olink

import (
	"testbed/testbed2/api"
)

type ManyParamInterfaceProperties struct {
	Prop1 *int32 `json:"prop1,omitempty"`
	Prop2 *int32 `json:"prop2,omitempty"`
	Prop3 *int32 `json:"prop3,omitempty"`
	Prop4 *int32 `json:"prop4,omitempty"`
}

// func ManyParamInterface.func1(param1 int32) int32

type ManyParamInterfaceFunc1Request struct {
	Param1 int32 `json:"param1"`
}

type ManyParamInterfaceFunc1Reply struct {
}

// func ManyParamInterface.func2(param1 int32, param2 int32) int32

type ManyParamInterfaceFunc2Request struct {
	Param1 int32 `json:"param1"`
	Param2 int32 `json:"param2"`
}

type ManyParamInterfaceFunc2Reply struct {
}

// func ManyParamInterface.func3(param1 int32, param2 int32, param3 int32) int32

type ManyParamInterfaceFunc3Request struct {
	Param1 int32 `json:"param1"`
	Param2 int32 `json:"param2"`
	Param3 int32 `json:"param3"`
}

type ManyParamInterfaceFunc3Reply struct {
}

// func ManyParamInterface.func4(param1 int32, param2 int32, param3 int32, param4 int32) int32

type ManyParamInterfaceFunc4Request struct {
	Param1 int32 `json:"param1"`
	Param2 int32 `json:"param2"`
	Param3 int32 `json:"param3"`
	Param4 int32 `json:"param4"`
}

type ManyParamInterfaceFunc4Reply struct {
}

type NestedStruct1InterfaceProperties struct {
	Prop1 *api.NestedStruct1 `json:"prop1,omitempty"`
}

// func NestedStruct1Interface.func1(param1 api.NestedStruct1) api.NestedStruct1

type NestedStruct1InterfaceFunc1Request struct {
	Param1 api.NestedStruct1 `json:"param1"`
}

type NestedStruct1InterfaceFunc1Reply struct {
}

type NestedStruct2InterfaceProperties struct {
	Prop1 *api.NestedStruct1 `json:"prop1,omitempty"`
	Prop2 *api.NestedStruct2 `json:"prop2,omitempty"`
}

// func NestedStruct2Interface.func1(param1 api.NestedStruct1) api.NestedStruct1

type NestedStruct2InterfaceFunc1Request struct {
	Param1 api.NestedStruct1 `json:"param1"`
}

type NestedStruct2InterfaceFunc1Reply struct {
}

// func NestedStruct2Interface.func2(param1 api.NestedStruct1, param2 api.NestedStruct2) api.NestedStruct1

type NestedStruct2InterfaceFunc2Request struct {
	Param1 api.NestedStruct1 `json:"param1"`
	Param2 api.NestedStruct2 `json:"param2"`
}

type NestedStruct2InterfaceFunc2Reply struct {
}

type NestedStruct3InterfaceProperties struct {
	Prop1 *api.NestedStruct1 `json:"prop1,omitempty"`
	Prop2 *api.NestedStruct2 `json:"prop2,omitempty"`
	Prop3 *api.NestedStruct3 `json:"prop3,omitempty"`
}

// func NestedStruct3Interface.func1(param1 api.NestedStruct1) api.NestedStruct1

type NestedStruct3InterfaceFunc1Request struct {
	Param1 api.NestedStruct1 `json:"param1"`
}

type NestedStruct3InterfaceFunc1Reply struct {
}

// func NestedStruct3Interface.func2(param1 api.NestedStruct1, param2 api.NestedStruct2) api.NestedStruct1

type NestedStruct3InterfaceFunc2Request struct {
	Param1 api.NestedStruct1 `json:"param1"`
	Param2 api.NestedStruct2 `json:"param2"`
}

type NestedStruct3InterfaceFunc2Reply struct {
}

// func NestedStruct3Interface.func3(param1 api.NestedStruct1, param2 api.NestedStruct2, param3 api.NestedStruct3) api.NestedStruct1

type NestedStruct3InterfaceFunc3Request struct {
	Param1 api.NestedStruct1 `json:"param1"`
	Param2 api.NestedStruct2 `json:"param2"`
	Param3 api.NestedStruct3 `json:"param3"`
}

type NestedStruct3InterfaceFunc3Reply struct {
}
