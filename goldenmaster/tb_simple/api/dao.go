package api

// func SimpleInterface.funcBool(paramBool bool) bool

type SimpleInterfaceFuncBoolRequest struct {
    ParamBool bool `json:"paramBool"`
}

type SimpleInterfaceFuncBoolReply struct {
    Result bool `json:"result"`
}

// func SimpleInterface.funcInt(paramInt int64) int64

type SimpleInterfaceFuncIntRequest struct {
    ParamInt int64 `json:"paramInt"`
}

type SimpleInterfaceFuncIntReply struct {
    Result int64 `json:"result"`
}

// func SimpleInterface.funcFloat(paramFloat float64) float64

type SimpleInterfaceFuncFloatRequest struct {
    ParamFloat float64 `json:"paramFloat"`
}

type SimpleInterfaceFuncFloatReply struct {
    Result float64 `json:"result"`
}

// func SimpleInterface.funcString(paramString string) string

type SimpleInterfaceFuncStringRequest struct {
    ParamString string `json:"paramString"`
}

type SimpleInterfaceFuncStringReply struct {
    Result string `json:"result"`
}

// func SimpleArrayInterface.funcBool(paramBool []bool) []bool

type SimpleArrayInterfaceFuncBoolRequest struct {
    ParamBool []bool `json:"paramBool"`
}

type SimpleArrayInterfaceFuncBoolReply struct {
    Result []bool `json:"result"`
}

// func SimpleArrayInterface.funcInt(paramInt []int64) []int64

type SimpleArrayInterfaceFuncIntRequest struct {
    ParamInt []int64 `json:"paramInt"`
}

type SimpleArrayInterfaceFuncIntReply struct {
    Result []int64 `json:"result"`
}

// func SimpleArrayInterface.funcFloat(paramFloat []float64) []float64

type SimpleArrayInterfaceFuncFloatRequest struct {
    ParamFloat []float64 `json:"paramFloat"`
}

type SimpleArrayInterfaceFuncFloatReply struct {
    Result []float64 `json:"result"`
}

// func SimpleArrayInterface.funcString(paramString []string) []string

type SimpleArrayInterfaceFuncStringRequest struct {
    ParamString []string `json:"paramString"`
}

type SimpleArrayInterfaceFuncStringReply struct {
    Result []string `json:"result"`
}

