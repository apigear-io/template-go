package olink


import (
    "goldenmaster/tb_again/api"
)

type SameStruct1InterfaceProperties struct {
    Prop1 *api.Struct1 `json:"prop1,omitempty"`
}

type SameStruct2InterfaceProperties struct {
    Prop1 *api.Struct2 `json:"prop1,omitempty"`
    Prop2 *api.Struct2 `json:"prop2,omitempty"`
}

type SameEnum1InterfaceProperties struct {
    Prop1 *api.Enum1 `json:"prop1,omitempty"`
}

type SameEnum2InterfaceProperties struct {
    Prop1 *api.Enum1 `json:"prop1,omitempty"`
    Prop2 *api.Enum2 `json:"prop2,omitempty"`
}

