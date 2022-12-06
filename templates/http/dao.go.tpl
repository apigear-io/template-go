package {{.Module.Name}}

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

{{- range .Module.Structs }}
{{- $class := .Name }}
type {{$class}} struct {
{{- range .Fields }}
    {{Camel .Name}} {{goReturn "" .}} `json:"{{camel .Name}}"`
{{- end }}
}
{{- end }}


{{- range .Module.Interfaces }}
{{- $iface := .Name }}

{{- range .Operations }}

{{- $name := (print $iface (Camel .Name))}}

// {{ snake $.Module.Name }}.{{ snake $iface}}/{{ snake .Name }}
type {{ $name }}Request struct {
{{- range .Params }}
    {{Camel .Name}} {{goReturn "" .}} `json:"{{camel .Name}}"`
{{- end }}
}

type {{ $name }}Reply struct {
{{- if .Return.IsVoid }}
    {{- if not .Return.IsVoid }}
    Result {{goReturn "" .Return}} `json:"result"`
    {{- end }}
{{- end }}
}

{{- end }}
{{- end }}

