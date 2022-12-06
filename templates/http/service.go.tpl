package {{.Module.Name}}

// Service Interfaces
{{- range .Module.Interfaces }}
{{- $iface := .Name }}
type I{{.Name}}Service interface {
{{- range .Operations }}
{{- $name := printf "%s%s" (Camel $iface) (Camel .Name) }}
    {{Camel .Name}}(req *{{$name}}Request) (*{{$name}}Reply, error)
{{- end }}
}
{{- end }}