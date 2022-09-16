package {{.Module.ShortName|lower}}

import "testing"

{{- range .Interface.Operations }}

func Test{{$.Interface.Name}}{{Camel .Name}}(t *testing.T) {
	t.Skip("TODO")
}
{{- end }}
