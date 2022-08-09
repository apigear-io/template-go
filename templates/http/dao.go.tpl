package http


import (
    "{{ .System.Name }}/{{snake .Module.Name}}/api"
)


{{- range .Module.Interfaces }}
{{- $iface := .Name }}

{{- range .Methods }}

{{- $name := (print $iface (Camel .Name))}}

// func {{ $iface }}.{{.Name}}({{goParams "api." .Inputs}}) {{goReturn "api." .Output}}

type {{ $name }}Request struct {
{{- range .Inputs }}
    {{camel .Name}} {{goReturn "api." .}} `json:"{{camel .Name}}"`
{{- end }}
}

type {{ $name }}Reply struct {
{{- if .Output.HasType }}
    Result {{goReturn "api." .Output}} `json:"result"`
{{- end }}
}

{{- end }}
{{- end }}

