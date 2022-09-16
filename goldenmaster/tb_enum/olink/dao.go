package olink


import (
    "goldenmaster/tb_enum/api"
)

type EnumInterfaceProperties struct {
    Prop0 *api.Enum0 `json:"prop0,omitempty"`
    Prop1 *api.Enum1 `json:"prop1,omitempty"`
    Prop2 *api.Enum2 `json:"prop2,omitempty"`
    Prop3 *api.Enum3 `json:"prop3,omitempty"`
}

// func EnumInterface.func0(param0 api.Enum0) api.Enum0

type EnumInterfaceFunc0Request struct {
    Param0 api.Enum0 `json:"param0"`
}

type EnumInterfaceFunc0Reply struct {
    Result api.Enum0 `json:"result"`
}

// func EnumInterface.func1(param1 api.Enum1) api.Enum1

type EnumInterfaceFunc1Request struct {
    Param1 api.Enum1 `json:"param1"`
}

type EnumInterfaceFunc1Reply struct {
    Result api.Enum1 `json:"result"`
}

// func EnumInterface.func2(param2 api.Enum2) api.Enum2

type EnumInterfaceFunc2Request struct {
    Param2 api.Enum2 `json:"param2"`
}

type EnumInterfaceFunc2Reply struct {
    Result api.Enum2 `json:"result"`
}

// func EnumInterface.func3(param3 api.Enum3) api.Enum3

type EnumInterfaceFunc3Request struct {
    Param3 api.Enum3 `json:"param3"`
}

type EnumInterfaceFunc3Reply struct {
    Result api.Enum3 `json:"result"`
}

