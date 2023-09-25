package olink

import (
	"testbed/testbed1/api"
)

type StructInterfaceProperties struct {
	PropBool   *api.StructBool   `json:"propBool,omitempty"`
	PropInt    *api.StructInt    `json:"propInt,omitempty"`
	PropFloat  *api.StructFloat  `json:"propFloat,omitempty"`
	PropString *api.StructString `json:"propString,omitempty"`
}

// func StructInterface.funcBool(paramBool api.StructBool) api.StructBool

type StructInterfaceFuncBoolRequest struct {
	ParamBool api.StructBool `json:"paramBool"`
}

type StructInterfaceFuncBoolReply struct {
}

// func StructInterface.funcInt(paramInt api.StructInt) api.StructInt

type StructInterfaceFuncIntRequest struct {
	ParamInt api.StructInt `json:"paramInt"`
}

type StructInterfaceFuncIntReply struct {
}

// func StructInterface.funcFloat(paramFloat api.StructFloat) api.StructFloat

type StructInterfaceFuncFloatRequest struct {
	ParamFloat api.StructFloat `json:"paramFloat"`
}

type StructInterfaceFuncFloatReply struct {
}

// func StructInterface.funcString(paramString api.StructString) api.StructString

type StructInterfaceFuncStringRequest struct {
	ParamString api.StructString `json:"paramString"`
}

type StructInterfaceFuncStringReply struct {
}

type StructArrayInterfaceProperties struct {
	PropBool   *[]api.StructBool   `json:"propBool,omitempty"`
	PropInt    *[]api.StructInt    `json:"propInt,omitempty"`
	PropFloat  *[]api.StructFloat  `json:"propFloat,omitempty"`
	PropString *[]api.StructString `json:"propString,omitempty"`
}

// func StructArrayInterface.funcBool(paramBool []api.StructBool) []api.StructBool

type StructArrayInterfaceFuncBoolRequest struct {
	ParamBool []api.StructBool `json:"paramBool"`
}

type StructArrayInterfaceFuncBoolReply struct {
}

// func StructArrayInterface.funcInt(paramInt []api.StructInt) []api.StructInt

type StructArrayInterfaceFuncIntRequest struct {
	ParamInt []api.StructInt `json:"paramInt"`
}

type StructArrayInterfaceFuncIntReply struct {
}

// func StructArrayInterface.funcFloat(paramFloat []api.StructFloat) []api.StructFloat

type StructArrayInterfaceFuncFloatRequest struct {
	ParamFloat []api.StructFloat `json:"paramFloat"`
}

type StructArrayInterfaceFuncFloatReply struct {
}

// func StructArrayInterface.funcString(paramString []api.StructString) []api.StructString

type StructArrayInterfaceFuncStringRequest struct {
	ParamString []api.StructString `json:"paramString"`
}

type StructArrayInterfaceFuncStringReply struct {
}
