package {{.Module.ShortName}}

import (
	"{{ .System.Name }}/{{ .Module.Name|snake }}/api"
)

{{- $self := .Interface.Name | first }}
{{- $class := printf "%sImpl" (camel .Interface.Name) }}

type {{$class}} struct {
    api.INotifier
{{- range .Interface.Properties }}
    {{.Name}} {{goReturn "api." . }}
{{- end }}
}

var _ api.{{.Interface.Name}} = (*{{$class}})(nil)
var _ api.INotifier = (*{{$class}})(nil)

func New{{.Interface.Name}}(notifier api.INotifier) api.{{.Interface.Name}} {
	obj := &{{$class}}{
        INotifier: notifier,
{{- range .Interface.Properties }}
{{- if .IsArray }}
        {{.Name}}: make({{goReturn "api." . }}, 0),
{{- else }}
        {{.Name}}: {{goDefault "api." . }},
{{- end}}
{{- end }}
    }
  	return obj
}

{{- range .Interface.Properties }}
// property get {{.Name}}
func ({{$self}} *{{$class}}) Get{{Camel .Name}}() {{goReturn "api." .}} {
	return {{goDefault "api." .}}
}

// property set {{.Name}}
func ({{$self}} *{{$class}}) Set{{Camel .Name}}({{goParam "api." .}}) {  
}
{{ end }}
{{- range .Interface.Operations }}
// method {{.Name}}
func ({{$self}} *{{$class}}) {{Camel .Name}}({{goParams "api." .Params}}) {{goReturn "api." .Return}} {
  return {{goDefault "api." .Return}}
}
    
{{ end }}
