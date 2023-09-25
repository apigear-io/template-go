package api

// func StructInterface.funcBool(paramBool StructBool) StructBool

type StructInterfaceFuncBoolRequest struct {
	ParamBool StructBool `json:"paramBool"`
}

type StructInterfaceFuncBoolReply struct {
}

// func StructInterface.funcInt(paramInt StructInt) StructInt

type StructInterfaceFuncIntRequest struct {
	ParamInt StructInt `json:"paramInt"`
}

type StructInterfaceFuncIntReply struct {
}

// func StructInterface.funcFloat(paramFloat StructFloat) StructFloat

type StructInterfaceFuncFloatRequest struct {
	ParamFloat StructFloat `json:"paramFloat"`
}

type StructInterfaceFuncFloatReply struct {
}

// func StructInterface.funcString(paramString StructString) StructString

type StructInterfaceFuncStringRequest struct {
	ParamString StructString `json:"paramString"`
}

type StructInterfaceFuncStringReply struct {
}

// func StructArrayInterface.funcBool(paramBool []StructBool) []StructBool

type StructArrayInterfaceFuncBoolRequest struct {
	ParamBool []StructBool `json:"paramBool"`
}

type StructArrayInterfaceFuncBoolReply struct {
}

// func StructArrayInterface.funcInt(paramInt []StructInt) []StructInt

type StructArrayInterfaceFuncIntRequest struct {
	ParamInt []StructInt `json:"paramInt"`
}

type StructArrayInterfaceFuncIntReply struct {
}

// func StructArrayInterface.funcFloat(paramFloat []StructFloat) []StructFloat

type StructArrayInterfaceFuncFloatRequest struct {
	ParamFloat []StructFloat `json:"paramFloat"`
}

type StructArrayInterfaceFuncFloatReply struct {
}

// func StructArrayInterface.funcString(paramString []StructString) []StructString

type StructArrayInterfaceFuncStringRequest struct {
	ParamString []StructString `json:"paramString"`
}

type StructArrayInterfaceFuncStringReply struct {
}
