package http


import (
    "{{ .System.Name }}/{{snake .Module.Name}}/api"
)


{{- range .Module.Interfaces }}
{{- $iface := .Name }}

{{- range .Operations }}

{{- $name := (print $iface (Camel .Name))}}

// func {{ $iface }}.{{.Name}}({{goParams "api." .Params}}) {{goReturn "api." .Return}}

type {{ $name }}Request struct {
{{- range .Params }}
    {{Camel .Name}} {{goReturn "api." .}} `json:"{{camel .Name}}"`
{{- end }}
}

type {{ $name }}Reply struct {
{{- if .Return.IsVoid }}
    Result {{goReturn "api." .Return}} `json:"result"`
{{- end }}
}

{{- end }}
{{- end }}

