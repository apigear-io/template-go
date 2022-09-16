{{- /*gotype: github.com/apigear-io/cli/pkg/model.ModuleScope*/ -}}
package api
{{ range .Module.Enums}}
{{- $class := .Name }}
// enum {{ $class }}
type {{$class}} int
const (
    {{$class}}_UNKNOWN {{$class}} = -1
{{- range .Members }}
    {{$class}}{{Camel .Name}} {{$class}} = {{.Value}}
{{- end }}
)
{{ end }}

// Data Structures
{{- range .Module.Structs }}
{{- $class := .Name }}
type {{$class}} struct {
{{- range .Fields }}
    {{Camel .Name}} {{goReturn "" .}} `json:"{{camel .Name}}"`
{{- end }}
}
{{- end }}

// Object Interfaces


type INotifier interface {
	NotifyPropertyChanged(name string, value any)
	NotifySignal(name string, args []any)
}

{{- range .Module.Interfaces }}
{{- $class := .Name }}
type {{$class}} interface {
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
{{- end }}

