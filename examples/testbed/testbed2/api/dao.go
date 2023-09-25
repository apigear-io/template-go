package api

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

// func NestedStruct1Interface.func1(param1 NestedStruct1) NestedStruct1

type NestedStruct1InterfaceFunc1Request struct {
	Param1 NestedStruct1 `json:"param1"`
}

type NestedStruct1InterfaceFunc1Reply struct {
}

// func NestedStruct2Interface.func1(param1 NestedStruct1) NestedStruct1

type NestedStruct2InterfaceFunc1Request struct {
	Param1 NestedStruct1 `json:"param1"`
}

type NestedStruct2InterfaceFunc1Reply struct {
}

// func NestedStruct2Interface.func2(param1 NestedStruct1, param2 NestedStruct2) NestedStruct1

type NestedStruct2InterfaceFunc2Request struct {
	Param1 NestedStruct1 `json:"param1"`
	Param2 NestedStruct2 `json:"param2"`
}

type NestedStruct2InterfaceFunc2Reply struct {
}

// func NestedStruct3Interface.func1(param1 NestedStruct1) NestedStruct1

type NestedStruct3InterfaceFunc1Request struct {
	Param1 NestedStruct1 `json:"param1"`
}

type NestedStruct3InterfaceFunc1Reply struct {
}

// func NestedStruct3Interface.func2(param1 NestedStruct1, param2 NestedStruct2) NestedStruct1

type NestedStruct3InterfaceFunc2Request struct {
	Param1 NestedStruct1 `json:"param1"`
	Param2 NestedStruct2 `json:"param2"`
}

type NestedStruct3InterfaceFunc2Reply struct {
}

// func NestedStruct3Interface.func3(param1 NestedStruct1, param2 NestedStruct2, param3 NestedStruct3) NestedStruct1

type NestedStruct3InterfaceFunc3Request struct {
	Param1 NestedStruct1 `json:"param1"`
	Param2 NestedStruct2 `json:"param2"`
	Param3 NestedStruct3 `json:"param3"`
}

type NestedStruct3InterfaceFunc3Reply struct {
}
