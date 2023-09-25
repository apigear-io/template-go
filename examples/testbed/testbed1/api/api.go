package api

// Data Structures
type StructBool struct {
	FieldBool bool `json:"fieldBool"`
}
type StructInt struct {
	FieldInt int32 `json:"fieldInt"`
}
type StructFloat struct {
	FieldFloat float32 `json:"fieldFloat"`
}
type StructString struct {
	FieldString string `json:"fieldString"`
}

// Object Interfaces

type INotifier interface {
	NotifyPropertyChanged(name string, value any)
	NotifySignal(name string, args []any)
}
type StructInterface interface {
	INotifier
	// Properties
	GetPropBool() StructBool
	SetPropBool(propBool StructBool)
	GetPropInt() StructInt
	SetPropInt(propInt StructInt)
	GetPropFloat() StructFloat
	SetPropFloat(propFloat StructFloat)
	GetPropString() StructString
	SetPropString(propString StructString)
	// Methods
	FuncBool(paramBool StructBool) StructBool
	FuncInt(paramInt StructInt) StructInt
	FuncFloat(paramFloat StructFloat) StructFloat
	FuncString(paramString StructString) StructString
}
type StructArrayInterface interface {
	INotifier
	// Properties
	GetPropBool() []StructBool
	SetPropBool(propBool []StructBool)
	GetPropInt() []StructInt
	SetPropInt(propInt []StructInt)
	GetPropFloat() []StructFloat
	SetPropFloat(propFloat []StructFloat)
	GetPropString() []StructString
	SetPropString(propString []StructString)
	// Methods
	FuncBool(paramBool []StructBool) []StructBool
	FuncInt(paramInt []StructInt) []StructInt
	FuncFloat(paramFloat []StructFloat) []StructFloat
	FuncString(paramString []StructString) []StructString
}
