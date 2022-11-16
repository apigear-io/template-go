package testbed

import (
    "fmt"
    "olink/pkg/client"
    "olink/pkg/ws"
    "{{ .System.Name }}/{{snake .Module.Name}}/api"
    "{{ .System.Name }}/{{snake .Module.Name}}/olink"
    "testing"
    
    "github.com/stretchr/testify/assert"
)

func getConnection(t *testing.T) (*client.Node, func(), error) {
	var addr = "ws://127.0.0.1:8080/ws"

	var registry = client.NewRegistry()
	var node = client.NewNode(registry)


	conn, err := ws.Dial(addr)
	if err != nil {
		return nil, nil, fmt.Errorf("dial error: %s\n", err)
	}
    done := func() {
        conn.Close()
    }
	node.SetOutput(conn)
	conn.SetOutput(node)
	registry.AttachClientNode(node)
    return node, done, nil
}

{{- range .Module.Interfaces }}

func Test{{Camel .Name}}(t *testing.T) {
    node, done, err := getConnection(t)
    defer done()
	assert.NoError(t, err)
    sink := olink.New{{Camel .Name}}Sink(node)
	node.LinkRemoteNode(sink.ObjectId())
	node.Registry.AddObjectSink(sink)
	node.LinkRemoteNode(sink.ObjectId())
    {{- range $i, $o := .Operations }}
    t.Run("{{.Name}}", func(t *testing.T) {
        {{- range .Params}}
        {{ camel .Name }} := {{ goDefault "api." .}}
        {{- end }}
        result{{Camel .Name}} := sink.{{Camel .Name}}({{ join ", " .ParamNames }})
        {{- if .Params }}
        {{- $p := index .Params 0 }}
        assert.Equal(t, {{camel $p.Name}}, result{{Camel .Name}})
        {{- end }}
        {{- range .Params}}
        {{- $prop := replace .Name "param" "prop"}}
        assert.Equal(t, {{.Name}}, sink.{{Camel $prop}})
        {{- end }}
    })
    {{- end }}
}
{{- end }}