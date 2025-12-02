package {{snake .Module.Name}}

{{- $self := .Interface.Name | first }}
{{- $class := printf "%sImpl" (camel .Interface.Name) }}

type {{$class}} struct {
    INotifier
{{- range .Interface.Properties }}
    {{.Name}} {{goReturn "" . }}
{{- end }}
}

var _ I{{.Interface.Name}} = (*{{$class}})(nil)
var _ INotifier = (*{{$class}})(nil)

func New{{.Interface.Name}}(notifier INotifier) I{{.Interface.Name}} {
    obj := &{{$class}}{
        INotifier: notifier,
        {{- range .Interface.Properties }}
        {{- if .IsArray }}
        {{.Name}}: make({{goReturn "" . }}, 0),
        {{- else }}
        {{.Name}}: {{goDefault "" . }},
        {{- end}}
        {{- end }}
    }
  	return obj
}

{{- range .Interface.Properties }}
// property get {{.Name}}
func ({{$self}} *{{$class}}) Get{{Camel .Name}}() {{goReturn "" .}} {
	return {{goDefault "" .}}
}

// property set {{camel .Name}}
func ({{$self}} *{{$class}}) Set{{Camel .Name}}({{goParam "" .}}) {
}
{{ end }}
{{- range .Interface.Operations }}
// method {{.Name}}
func ({{$self}} *{{$class}}) {{Camel .Name}}({{goParams "" .Params}}) {{goReturn "" .Return}} {
  return {{goDefault "" .Return}}
}

{{ end }}
