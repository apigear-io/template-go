{{- /*gotype: github.com/apigear-io/cli/pkg/model.ModuleScope*/ -}}
package api


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
{{- if .Return.IsVoid }}
    Result {{goReturn "" .Return}} `json:"result"`
{{- end }}
}

{{- end }}
{{- end }}

