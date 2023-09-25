package api

// enum Conflict8
type Conflict8 int

const (
	Conflict8_UNKNOWN Conflict8 = -1
)

// enum Conflict9
type Conflict9 int

const (
	Conflict9_UNKNOWN Conflict9 = -1
)

// enum Conflict10
type Conflict10 int

const (
	Conflict10_UNKNOWN Conflict10 = -1
)

// Data Structures
type Conflict5 struct {
}
type Conflict6 struct {
}
type Conflict7 struct {
}

// Object Interfaces

type INotifier interface {
	NotifyPropertyChanged(name string, value any)
	NotifySignal(name string, args []any)
}
type Conflict1 interface {
	INotifier
	// Properties
	GetSameName() int64
	SetSameName(sameName int64)
	// Methods
}
type Conflict2 interface {
	INotifier
	// Properties
	GetSameName() int64
	SetSameName(sameName int64)
	// Methods
}
type Conflict3 interface {
	INotifier
	// Properties
	// Methods
}
type Conflict4 interface {
	INotifier
	// Properties
	// Methods
}
