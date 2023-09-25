package olink

import (
	"fmt"
	"log"
	"testbed/tb_simple/api"

	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
)

type SimpleInterfaceSource struct {
	node *remote.Node
	impl api.SimpleInterface
}

var _ remote.IObjectSource = (*SimpleInterfaceSource)(nil)
var _ api.INotifier = (*SimpleInterfaceSource)(nil)

func NewSimpleInterfaceSource() *SimpleInterfaceSource {
	return &SimpleInterfaceSource{}
}

func (s *SimpleInterfaceSource) SetImplementation(impl api.SimpleInterface) {
	s.impl = impl
}

func (s *SimpleInterfaceSource) ObjectId() string {
	return "tb.simple.SimpleInterface"
}

func (s *SimpleInterfaceSource) Invoke(methodId string, args core.Args) (core.Any, error) {
	log.Printf("source: invoke: %s %v", methodId, args)
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	name := core.SymbolIdToMember(methodId)
	switch name {
	case "funcBool":
		paramBool, err := api.AsBool(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.FuncBool(paramBool), nil
	case "funcInt":
		paramInt, err := api.AsInt(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.FuncInt(paramInt), nil
	case "funcFloat":
		paramFloat, err := api.AsFloat(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.FuncFloat(paramFloat), nil
	case "funcString":
		paramString, err := api.AsString(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.FuncString(paramString), nil
	default:
		return nil, fmt.Errorf("unknown method: %s", name)
	}
}

func (s *SimpleInterfaceSource) SetProperty(propertyId string, value core.Any) error {
	if s.impl == nil {
		return fmt.Errorf("no implementation")
	}
	name := core.SymbolIdToMember(propertyId)
	switch name {
	case "propBool":
		prop, err := api.AsBool(value)
		if err != nil {
			return err
		}
		s.impl.SetPropBool(prop)
	case "propInt":
		prop, err := api.AsInt(value)
		if err != nil {
			return err
		}
		s.impl.SetPropInt(prop)
	case "propFloat":
		prop, err := api.AsFloat(value)
		if err != nil {
			return err
		}
		s.impl.SetPropFloat(prop)
	case "propString":
		prop, err := api.AsString(value)
		if err != nil {
			return err
		}
		s.impl.SetPropString(prop)
	default:
		return fmt.Errorf("SimpleInterfaceSource.SetProperty: unknown property %s", propertyId)
	}
	return nil
}

func (s *SimpleInterfaceSource) CollectProperties() (core.KWArgs, error) {
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	return core.KWArgs{
		"propBool":   s.impl.GetPropBool(),
		"propInt":    s.impl.GetPropInt(),
		"propFloat":  s.impl.GetPropFloat(),
		"propString": s.impl.GetPropString(),
	}, nil
}

func (s *SimpleInterfaceSource) Linked(objectId string, node *remote.Node) error {
	if objectId != s.ObjectId() {
		return fmt.Errorf("not my object: %s", objectId)
	}
	s.node = node
	return nil
}

func (s *SimpleInterfaceSource) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *SimpleInterfaceSource) NotifySignal(name string, args []any) {
	signalId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}
