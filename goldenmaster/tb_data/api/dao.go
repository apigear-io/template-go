package api

// func StructInterface.funcBool(paramBool StructBool) StructBool

type StructInterfaceFuncBoolRequest struct {
	ParamBool StructBool `json:"paramBool"`
}

type StructInterfaceFuncBoolReply struct {
	Result StructBool `json:"result"`
}

// func StructInterface.funcInt(paramInt StructInt) StructInt

type StructInterfaceFuncIntRequest struct {
	ParamInt StructInt `json:"paramInt"`
}

type StructInterfaceFuncIntReply struct {
	Result StructInt `json:"result"`
}

// func StructInterface.funcFloat(paramFloat StructFloat) StructFloat

type StructInterfaceFuncFloatRequest struct {
	ParamFloat StructFloat `json:"paramFloat"`
}

type StructInterfaceFuncFloatReply struct {
	Result StructFloat `json:"result"`
}

// func StructInterface.funcString(paramString StructString) StructString

type StructInterfaceFuncStringRequest struct {
	ParamString StructString `json:"paramString"`
}

type StructInterfaceFuncStringReply struct {
	Result StructString `json:"result"`
}

// func StructArrayInterface.funcBool(paramBool []StructBool) []StructBool

type StructArrayInterfaceFuncBoolRequest struct {
	ParamBool []StructBool `json:"paramBool"`
}

type StructArrayInterfaceFuncBoolReply struct {
	Result []StructBool `json:"result"`
}

// func StructArrayInterface.funcInt(paramInt []StructInt) []StructInt

type StructArrayInterfaceFuncIntRequest struct {
	ParamInt []StructInt `json:"paramInt"`
}

type StructArrayInterfaceFuncIntReply struct {
	Result []StructInt `json:"result"`
}

// func StructArrayInterface.funcFloat(paramFloat []StructFloat) []StructFloat

type StructArrayInterfaceFuncFloatRequest struct {
	ParamFloat []StructFloat `json:"paramFloat"`
}

type StructArrayInterfaceFuncFloatReply struct {
	Result []StructFloat `json:"result"`
}

// func StructArrayInterface.funcString(paramString []StructString) []StructString

type StructArrayInterfaceFuncStringRequest struct {
	ParamString []StructString `json:"paramString"`
}

type StructArrayInterfaceFuncStringReply struct {
	Result []StructString `json:"result"`
}
