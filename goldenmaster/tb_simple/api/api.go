package api


// Data Structures

// Object Interfaces


type INotifier interface {
	NotifyPropertyChanged(name string, value any)
	NotifySignal(name string, args []any)
}
type SimpleInterface interface {
    INotifier
    // Properties
    GetPropBool() bool
    SetPropBool(propBool bool)
    GetPropInt() int64
    SetPropInt(propInt int64)
    GetPropFloat() float64
    SetPropFloat(propFloat float64)
    GetPropString() string
    SetPropString(propString string)
    // Methods
    FuncBool(paramBool bool) bool
    FuncInt(paramInt int64) int64
    FuncFloat(paramFloat float64) float64
    FuncString(paramString string) string
}
type SimpleArrayInterface interface {
    INotifier
    // Properties
    GetPropBool() []bool
    SetPropBool(propBool []bool)
    GetPropInt() []int64
    SetPropInt(propInt []int64)
    GetPropFloat() []float64
    SetPropFloat(propFloat []float64)
    GetPropString() []string
    SetPropString(propString []string)
    // Methods
    FuncBool(paramBool []bool) []bool
    FuncInt(paramInt []int64) []int64
    FuncFloat(paramFloat []float64) []float64
    FuncString(paramString []string) []string
}

