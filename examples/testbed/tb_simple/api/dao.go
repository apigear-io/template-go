package api

// func SimpleInterface.funcBool(paramBool bool) bool

type SimpleInterfaceFuncBoolRequest struct {
	ParamBool bool `json:"paramBool"`
}

type SimpleInterfaceFuncBoolReply struct {
}

// func SimpleInterface.funcInt(paramInt int32) int32

type SimpleInterfaceFuncIntRequest struct {
	ParamInt int32 `json:"paramInt"`
}

type SimpleInterfaceFuncIntReply struct {
}

// func SimpleInterface.funcFloat(paramFloat float32) float32

type SimpleInterfaceFuncFloatRequest struct {
	ParamFloat float32 `json:"paramFloat"`
}

type SimpleInterfaceFuncFloatReply struct {
}

// func SimpleInterface.funcString(paramString string) string

type SimpleInterfaceFuncStringRequest struct {
	ParamString string `json:"paramString"`
}

type SimpleInterfaceFuncStringReply struct {
}

// func SimpleArrayInterface.funcBool(paramBool []bool) []bool

type SimpleArrayInterfaceFuncBoolRequest struct {
	ParamBool []bool `json:"paramBool"`
}

type SimpleArrayInterfaceFuncBoolReply struct {
}

// func SimpleArrayInterface.funcInt(paramInt []int32) []int32

type SimpleArrayInterfaceFuncIntRequest struct {
	ParamInt []int32 `json:"paramInt"`
}

type SimpleArrayInterfaceFuncIntReply struct {
}

// func SimpleArrayInterface.funcFloat(paramFloat []float32) []float32

type SimpleArrayInterfaceFuncFloatRequest struct {
	ParamFloat []float32 `json:"paramFloat"`
}

type SimpleArrayInterfaceFuncFloatReply struct {
}

// func SimpleArrayInterface.funcString(paramString []string) []string

type SimpleArrayInterfaceFuncStringRequest struct {
	ParamString []string `json:"paramString"`
}

type SimpleArrayInterfaceFuncStringReply struct {
}
