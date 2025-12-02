package {{snake .Module.Name}}

{{- $import := .System.Meta.GetString "go.module" }}

import (
    "fmt"
    "sync"

    "github.com/apigear-io/objectlink-core-go/olink/client"
	"github.com/apigear-io/objectlink-core-go/olink/core"

	"github.com/rs/zerolog/log"
    "github.com/google/go-cmp/cmp"
)

{{- $class := printf "%sSink" .Interface }}

type {{$class}} struct {
    node *client.Node
{{- range .Interface.Properties }}
    {{Camel .Name }} {{ goReturn "" . }}
{{- end }}
{{- range .Interface.Signals }}
    On{{Camel .Name}} func({{goParams "" .Params}})
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



{{- range .Interface.Properties }}

func (s *{{$class}}) Set{{Camel .Name}}(v {{goReturn "" .}}) {
	if s.node == nil {
        return
    }
    if cmp.Equal(s.{{Camel .Name}}, v) {
        return
    }
    s.{{Camel .Name}} = v
    propertyId := core.MakeSymbolId(s.ObjectId(), "{{.Name}}")
    s.node.SetRemoteProperty(propertyId, v)
}
{{- end }}
{{- range .Interface.Operations }}

func (s *{{$class}}) {{Camel .Name}}({{ goParams "" .Params }}) {{ goReturn "" .Return }} {
    {{- if not .Return.IsVoid }}
    var reply {{goReturn "" .Return}}
    {{- end }}
    if s.node == nil {
        return
    }
    methodId := core.MakeSymbolId(s.ObjectId(), "{{.Name}}")
    {{- if .Return.IsVoid }}
    wg := sync.WaitGroup{}
    wg.Add(1)
    args := core.Args{ {{ range $i, $p := .Params }}{{- if $i }}, {{ end -}}{{- camel $p.Name -}} {{ end }} }
    s.node.InvokeRemote(methodId, args, func(arg client.InvokeReplyArg) {
        wg.Done()
        {{- if not .Return.IsVoid }}
        reply = arg.Value.({{goReturn "" .Return}})
        {{- end }}
    })
    wg.Wait()
    {{- else }}
    s.node.InvokeRemote(methodId, core.Args{}, nil)
    {{- end }}
    {{- if not .Return.IsVoid }}
    return reply
    {{- end }}
}
{{- end }}


func (s *{{$class}}) HandleSignal(signalId string, args core.Args) {
	log.Info().Msgf("handle signal: %s %v\n", signalId, args)
    member := core.SymbolIdToMember(signalId)
    switch member {
    {{- range .Interface.Signals }}
    case "{{.Name}}":
        if s.On{{Camel .Name}} != nil {
            s.On{{Camel .Name}}({{ goVars "" .Params }})
        }
    {{- end }}
    default:
        log.Warn().Msgf("unknown signal: %s", signalId) 
    }
}


func (s *{{$class}}) HandleInit(objectId string, props core.KWArgs, node *client.Node) {
	log.Info().Msgf("handle init: %s props = %#v\n", objectId, props)
	if objectId != s.ObjectId() {
        return
    }
    s.node = node
    {{- range .Interface.Properties }}
    if value, ok := props["{{.Name}}"]; ok {
        var v {{goReturn "" .}}
        if ConvertTo(value, &v) {            
            s.Set{{Camel .Name}}(v)
        }
    }
    {{- end }}
}

func (s *{{$class}}) HandlePropertyChange(propertyId string, value core.Any) {
	fmt.Printf("handle property change: %s %v\n", propertyId, value)
    member := core.SymbolIdToMember(propertyId)
    switch member {
    {{- range .Interface.Properties }}
    case "{{.Name}}":
        var v {{goReturn "" .}}
        if ConvertTo(value, &v) {
            s.Set{{Camel .Name}}(v)
        }
    {{- end }}
    default:
        log.Warn().Msgf("unknown property: %s", propertyId)
    }
}

func (s *{{$class}}) HandleRelease() {
	fmt.Printf("handle release: %s\n", s.ObjectId())
	if s.node != nil {
		s.node = nil
	}
}
