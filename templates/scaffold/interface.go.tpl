package {{.Module.ShortName}}

import (
	"{{ .System.Name }}/{{ .Module.Name|path }}/api"
)

{{- $self := .Interface.Name | first }}
{{- $class := .Interface.Name | camel }}

type {{$class}} struct {
{{- range .Interface.Properties }}
    {{.Name}} {{goReturn "api." . }}
{{- end }}
}

func New{{.Interface.Name}}() api.{{.Interface.Name}} {
  	obj := &{{$class}}{
{{- range .Interface.Properties }}
        {{.Name}}: {{goDefault "api." .}},
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
{{- range .Interface.Methods }}
// method {{.Name}}
func ({{$self}} *{{$class}}) {{.Name|camel}}({{goParams "api." .Inputs}}) {{goReturn "api." .Output}} {
  return {{goDefault "api." .Output}}
}
    
{{ end }}
