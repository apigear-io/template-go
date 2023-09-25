package olink

import (
    "fmt"
    "log"
	"github.com/apigear-io/objectlink-core-go/olink/core"
    "github.com/apigear-io/objectlink-core-go/olink/remote"
    "{{ .System.Name }}/{{snake .Module.Name}}/api"
)

{{- $class := printf "%sSource" .Interface.Name}}

type {{ $class }} struct {
    node *remote.Node
    impl api.{{ .Interface.Name }}
}

var _ remote.IObjectSource = (*{{$class}})(nil)
var _ api.INotifier = (*{{$class}})(nil)

func New{{$class}}() *{{ $class }} {
    return &{{ $class }}{}
}

func (s *{{$class}}) SetImplementation(impl api.{{ .Interface.Name }}) {
    s.impl = impl
}

func (s *{{$class}}) ObjectId() string {
	return "{{.Module.Name}}.{{.Interface.Name}}"
}

func (s *{{$class}}) Invoke(methodId string, args core.Args) (core.Any, error) {
	log.Printf("source: invoke: %s %v", methodId, args)
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	name := core.SymbolIdToMember(methodId)
	switch name {
{{- range .Interface.Operations }}
	case "{{.Name}}":
        {{- range $i, $p := .Params }}
        {{.Name}}, err := api.As{{ Camel .TypeName }}(args[{{$i}}])
        if err != nil {
            return nil, err
        }
        {{- end }}
        {{- if .Return.IsVoid }}
        s.impl.{{ Camel .Name}}({{join ", " .ParamNames}})
        return nil, nil
        {{- else }}
        return s.impl.{{ Camel .Name}}({{join ", " .ParamNames}}), nil
        {{- end }}
{{- end }}
    default:
        return nil, fmt.Errorf("unknown method: %s", name)
    }
}

func (s *{{$class}}) SetProperty(propertyId string, value core.Any) error {
    if s.impl == nil {
        return fmt.Errorf("no implementation")
    }
    name := core.SymbolIdToMember(propertyId)
    switch name {
    {{- range .Interface.Properties }}
    case "{{ .Name }}":
        prop, err := api.As{{ Camel .TypeName }}(value)
        if err != nil {
            return err
        }
        s.impl.Set{{ Camel .Name }}(prop)
    {{- end }}
    default:
        return fmt.Errorf("{{$class}}.SetProperty: unknown property %s", propertyId)
    }
    return nil
}

func (s *{{$class}}) CollectProperties() (core.KWArgs, error) {
    if s.impl == nil {
        return nil, fmt.Errorf("no implementation")
    }
    return core.KWArgs{
        {{- range .Interface.Properties }}
        "{{ .Name }}": s.impl.Get{{ Camel .Name }}(),
        {{- end }}
    }, nil
}


func (s *{{$class}}) Linked(objectId string, node *remote.Node) error {
    if objectId != s.ObjectId() {
        return fmt.Errorf("not my object: %s", objectId)
    }
    s.node = node
    return nil
}

func (s *{{$class}}) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *{{$class}}) NotifySignal(name string, args []any) {
	signalId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}
