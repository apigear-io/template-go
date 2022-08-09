package {{.Module.ShortName|lower}}

import "testing"

{{- range .Interface.Methods }}

func Test{{$.Interface.Name}}{{Camel .Name}}(t *testing.T) {
	t.Errorf("TODO")
}
{{- end }}
