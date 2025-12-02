package {{snake .Module.Name}}

import (
    "fmt"
	"github.com/go-viper/mapstructure/v2"
    "github.com/rs/zerolog/log"
    "github.com/google/go-cmp/cmp"
)

{{- range .Module.Enums}}
{{- $class := .Name }}
// enum {{ $class }}
type {{$class}} int
const (
    {{$class}}_UNKNOWN {{$class}} = -1
{{- range .Members }}
    {{$class}}{{Camel .Name}} {{$class}} = {{.Value}}
{{- end }}
)
{{- end }}

// Data Structures
{{- range .Module.Structs }}
{{- $class := .Name }}
type {{$class}} struct {
{{- range .Fields }}
    {{Camel .Name}} {{goReturn "" .}} `json:"{{camel .Name}}"`
{{- end }}
}
{{- end }}

// Notifier Interfaces
type INotifier interface {
	NotifyPropertyChanged(name string, value any)
	NotifySignal(name string, args []any)
}


func ConvertTo(v any, target any) bool {
	err := mapstructure.Decode(v, target)
	if err != nil {
		log.Error().Msgf("error: %v", err)
		return false
	}
	return true
}

{{- range .Module.Interfaces }}
{{- $class := (Camel .Name) }}

// {{ $class }} interface
type I{{$class}} interface {
    INotifier

    // Properties
{{- range .Properties }}
    Get{{Camel .Name}}() {{goReturn "" .}}
    Set{{Camel .Name}}({{goParam "" .}})
{{- end }}
    // Methods
{{- range .Operations }}
    {{Camel .Name}}({{goParams "" .Params}}) {{goReturn "" .Return}}
{{- end }}
}


var _ I{{$class}} = (*Default{{$class}})(nil)
var _ INotifier = (*Default{{$class}})(nil)

// Default Implementation
type Default{{$class}} struct {    
    notifier INotifier
    // Properties
{{- range .Properties }}
    {{camel .Name}} {{goReturn "" .}}
{{- end }}
}

func NewDefault{{$class}}(notifier INotifier) *Default{{$class}} {
    return &Default{{$class}}{
        notifier: notifier,
    }
}

func (d *Default{{$class}}) NotifyPropertyChanged(name string, value any) {
    if d.notifier != nil {
        d.notifier.NotifyPropertyChanged(name, value)
    }
}

func (d *Default{{$class}}) NotifySignal(name string, args []any) {
    if d.notifier != nil {
        d.notifier.NotifySignal(name, args)
    }
}


{{- range .Properties }}

// property {{.Name}}
func (d *Default{{$class}}) Get{{Camel .Name}}() {{goReturn "" .}} {
    return d.{{camel .Name}}
}

func (d *Default{{$class}}) Set{{Camel .Name}}(v {{goType "" .}}) {
    log.Info().Msgf("Setting {{camel .Name}} to %v", v)
    if cmp.Equal(d.{{camel .Name}}, v) {
        return
    }
    d.{{camel .Name}} = v
    d.NotifyPropertyChanged("{{camel .Name}}", v)
}

{{- end }}

{{- range .Operations }}

// operation {{.Name}}
func (d *Default{{$class}}) {{Camel .Name}}({{goParams "" .Params}}) {{goReturn "" .Return}} {
    fmt.Printf("Default{{$class}}: {{.Name}}(%v)\n", []any{ {{goVars .Params}} })
{{- if not .Return.IsVoid }}
    return {{goReturn "" .Return}}
{{- end }}
}
{{- end }}

{{- range .Signals }}

// signal {{.Name}}
func (d *Default{{$class}}) Notify{{Camel .Name}}({{goParams "" .Params}}) {
    d.NotifySignal("{{camel .Name}}", []any{ {{goVars .Params}} })
}
{{- end }}
{{- end }}

// Request and Reply Structures
{{- range .Module.Interfaces }}
{{- $iface := .Name }}

{{- range .Operations }}

{{- $name := (print $iface (Camel .Name))}}

// func {{ $iface }}.{{.Name}}({{goParams "" .Params}}) {{goReturn "" .Return}}

type {{ $name }}Request struct {
{{- range .Params }}
    {{Camel .Name}} {{goReturn "" .}} `json:"{{camel .Name}}"`
{{- end }}
}

type {{ $name }}Reply struct {
{{- if not .Return.IsVoid }}
    Result {{goReturn "" .Return}} `json:"result"`
{{- end }}
}

{{- end }}
{{- end }}
