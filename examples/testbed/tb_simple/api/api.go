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
	GetPropInt() int32
	SetPropInt(propInt int32)
	GetPropFloat() float32
	SetPropFloat(propFloat float32)
	GetPropString() string
	SetPropString(propString string)
	// Methods
	FuncBool(paramBool bool) bool
	FuncInt(paramInt int32) int32
	FuncFloat(paramFloat float32) float32
	FuncString(paramString string) string
}
type SimpleArrayInterface interface {
	INotifier
	// Properties
	GetPropBool() []bool
	SetPropBool(propBool []bool)
	GetPropInt() []int32
	SetPropInt(propInt []int32)
	GetPropFloat() []float32
	SetPropFloat(propFloat []float32)
	GetPropString() []string
	SetPropString(propString []string)
	// Methods
	FuncBool(paramBool []bool) []bool
	FuncInt(paramInt []int32) []int32
	FuncFloat(paramFloat []float32) []float32
	FuncString(paramString []string) []string
}
