package olink


import (
    "goldenmaster/tb_adv/api"
)

type ManyParamInterfaceProperties struct {
    Prop1 *int64 `json:"prop1,omitempty"`
    Prop2 *int64 `json:"prop2,omitempty"`
    Prop3 *int64 `json:"prop3,omitempty"`
    Prop4 *int64 `json:"prop4,omitempty"`
}

type NestedStruct1InterfaceProperties struct {
    Prop1 *api.NestedStruct1 `json:"prop1,omitempty"`
}

type NestedStruct2InterfaceProperties struct {
    Prop1 *api.NestedStruct1 `json:"prop1,omitempty"`
    Prop2 *api.NestedStruct2 `json:"prop2,omitempty"`
}

type NestedStruct3InterfaceProperties struct {
    Prop1 *api.NestedStruct1 `json:"prop1,omitempty"`
    Prop2 *api.NestedStruct2 `json:"prop2,omitempty"`
    Prop3 *api.NestedStruct3 `json:"prop3,omitempty"`
}

