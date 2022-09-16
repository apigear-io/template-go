package testbed

import (
	"{{ .System.Name }}/{{ .Module.Name|snake }}/api"
)

{{- $self := .Interface.Name | first }}
{{- $class := printf "%sImpl" (camel .Interface.Name) }}

type {{$class}} struct {
    api.INotifier
{{- range .Interface.Properties }}
    {{camel .Name}} {{goReturn "api." . }}
{{- end }}
}

var _ api.{{.Interface.Name}} = (*{{$class}})(nil)
var _ api.INotifier = (*{{$class}})(nil)

func New{{.Interface.Name}}(notifier api.INotifier) api.{{.Interface.Name}} {
	obj := &{{$class}}{
        INotifier: notifier,
{{- range .Interface.Properties }}
{{- if .IsArray }}
        {{camel .Name}}: make({{goReturn "api." . }}, 0),
{{- else }}
        {{camel .Name}}: {{goDefault "api." . }},
{{- end}}
{{- end }}
    }
  	return obj
}

{{- range .Interface.Properties }}
// property get {{.Name}}
func ({{$self}} *{{$class}}) Get{{Camel .Name}}() {{goReturn "api." .}} {
	return {{$self}}.{{camel .Name}}
}

// property set {{.Name}}
func ({{$self}} *{{$class}}) Set{{Camel .Name}}({{goParam "api." .}}) {
    {{$self}}.{{camel .Name}} = {{camel .Name}}
    {{$self}}.NotifyPropertyChanged("{{.Name}}", {{camel .Name}})
}
{{ end }}
{{- range .Interface.Operations }}
// method {{.Name}}
func ({{$self}} *{{$class}}) {{Camel .Name}}({{goParams "api." .Params}}) {{goReturn "api." .Return}} {
    {{- range .Params}}
    {{- $prop := replace .Name "param" "prop"}}
    {{$self}}.Set{{Camel $prop}}({{camel .Name}})
    {{- $sig := replace .Name "param" "sig"}}
    {{$self}}.NotifySignal("{{$sig}}", []any{ {{camel .Name}} })
    {{- end}}
    {{- if .Return.HasType }}
    {{- $p := index .Params 0 }}
    return {{$p.Name}}
    {{- end }}
}
    
{{ end }}
