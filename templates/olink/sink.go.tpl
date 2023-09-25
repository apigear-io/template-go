package olink

import (
    "fmt"
    "github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/core"
    "{{ .System.Name }}/{{snake .Module.Name}}/api"
)

{{ $class := printf "%sSink" .Interface }}
type {{$class}} struct {
    node *client.Node
{{- range .Interface.Properties }}
    {{Camel .Name }} {{ goReturn "api." . }} `json:"{{ .Name }}"`
{{- end }}
{{- range .Interface.Signals }}
    On{{Camel .Name}} func({{goParams "api." .Params}})
{{- end }}    
}

var _ client.IObjectSink = (*{{$class}})(nil)


func New{{$class}}(node *client.Node) *{{$class}} {
	return &{{$class}}{
		node: node,
	}
}

func (s *{{$class}}) ObjectId() string {
	return "{{dot .Module.Name}}.{{.Interface.Name}}"
}

{{ range .Interface.Properties }}
func (s *{{$class}}) Set{{Camel .Name}}({{goParam "api." .}}) {
	if s.node == nil {
        return
    }
    propertyId := core.MakeSymbolId(s.ObjectId(), "{{.Name}}")
    s.node.SetRemoteProperty(propertyId, {{.Name}})
}
{{ end }}
{{- range .Interface.Operations }}
func (s *{{$class}}) {{Camel .Name}}({{ goParams "api." .Params }}) {{ goReturn "api." .Return }} {
    {{- if not .Return.IsVoid }}
    var reply {{goReturn "api." .Return}}    
    {{- end }}    
    if s.node != nil {
        methodId := core.MakeSymbolId(s.ObjectId(), "{{.Name}}")
        {{- if .Return.IsVoid }}
        wg := sync.WaitGroup{}
        wg.Add(1)
        args := core.Args{ {{join ", " .ParamNames}} }
        s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
            wg.Done()
            reply = arg.Value.({{goReturn "api." .Return}})
        })
        wg.Wait()
        {{- else }}
        s.node.InvokeRemote(methodId, core.Args{}, nil)
        {{- end }}
    }
    {{- if not .Return.IsVoid }}
    return reply
    {{- end }}
}
{{ end }}


func (s *{{$class}}) OnSignal(signalId string, args core.Args) {
	fmt.Printf("on signal: %s %v\n", signalId, args)
    name := core.SymbolIdToMember(signalId)
    switch name {
    {{- range .Interface.Signals }}
    case "{{.Name}}":
        if s.On{{Camel .Name}} != nil {
            {{- range $i, $p := .Params }}
            {{.Name}} := args[{{$i}}].({{goReturn "api." .}})
            {{- end }}
            s.On{{Camel .Name}}(
                {{- range $i, $p :=  .Params -}}
                {{if $i}}, {{end}}{{.Name}}
                {{- end -}})
        }
    {{- end }}
    }
}


func (s *{{$class}}) OnInit(objectId string, props core.KWArgs, node *client.Node) {
	fmt.Printf("on init: %s props = %#v\n", objectId, props)
	if objectId != s.ObjectId() {
        return
    }
    s.node = node
    {{- range .Interface.Properties }}
    if value, ok := props["{{.Name}}"]; ok {
        v, err := api.As{{Camel .TypeName}}(value)
        if err != nil {
            fmt.Printf("error: %v\n", err)
        } else {
            s.Set{{Camel .Name}}(v)
        }
    }
    {{- end }}
}

func (s *{{$class}}) OnPropertyChange(propertyId string, value core.Any) {
	fmt.Printf("on property change: %s %v\n", propertyId, value)
	name := core.SymbolIdToMember(propertyId)
	switch name {
    {{- range .Interface.Properties }}
	case "{{.Name}}":
        v, err := api.As{{Camel .TypeName}}(value)
        if err != nil {
            fmt.Printf("error: %v\n", err)
        } else {
            s.Set{{Camel .Name}}(v)
        }
    {{- end }}
	default:
		fmt.Printf("unknown property: %s\n", propertyId)
	}
}


func (s *{{$class}}) OnRelease() {
	fmt.Printf("on release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}