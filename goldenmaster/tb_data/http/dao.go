package http

import (
	"goldenmaster/tb_data/api"
)

// func StructInterface.funcBool(paramBool api.StructBool) api.StructBool

type StructInterfaceFuncBoolRequest struct {
	ParamBool api.StructBool `json:"paramBool"`
}

type StructInterfaceFuncBoolReply struct {
	Result api.StructBool `json:"result"`
}

// func StructInterface.funcInt(paramInt api.StructInt) api.StructInt

type StructInterfaceFuncIntRequest struct {
	ParamInt api.StructInt `json:"paramInt"`
}

type StructInterfaceFuncIntReply struct {
	Result api.StructInt `json:"result"`
}

// func StructInterface.funcFloat(paramFloat api.StructFloat) api.StructFloat

type StructInterfaceFuncFloatRequest struct {
	ParamFloat api.StructFloat `json:"paramFloat"`
}

type StructInterfaceFuncFloatReply struct {
	Result api.StructFloat `json:"result"`
}

// func StructInterface.funcString(paramString api.StructString) api.StructString

type StructInterfaceFuncStringRequest struct {
	ParamString api.StructString `json:"paramString"`
}

type StructInterfaceFuncStringReply struct {
	Result api.StructString `json:"result"`
}

// func StructArrayInterface.funcBool(paramBool []api.StructBool) []api.StructBool

type StructArrayInterfaceFuncBoolRequest struct {
	ParamBool []api.StructBool `json:"paramBool"`
}

type StructArrayInterfaceFuncBoolReply struct {
	Result []api.StructBool `json:"result"`
}

// func StructArrayInterface.funcInt(paramInt []api.StructInt) []api.StructInt

type StructArrayInterfaceFuncIntRequest struct {
	ParamInt []api.StructInt `json:"paramInt"`
}

type StructArrayInterfaceFuncIntReply struct {
	Result []api.StructInt `json:"result"`
}

// func StructArrayInterface.funcFloat(paramFloat []api.StructFloat) []api.StructFloat

type StructArrayInterfaceFuncFloatRequest struct {
	ParamFloat []api.StructFloat `json:"paramFloat"`
}

type StructArrayInterfaceFuncFloatReply struct {
	Result []api.StructFloat `json:"result"`
}

// func StructArrayInterface.funcString(paramString []api.StructString) []api.StructString

type StructArrayInterfaceFuncStringRequest struct {
	ParamString []api.StructString `json:"paramString"`
}

type StructArrayInterfaceFuncStringReply struct {
	Result []api.StructString `json:"result"`
}
