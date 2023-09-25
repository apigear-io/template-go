package olink

import (
	"fmt"
	"log"
	"testbed/testbed1/api"

	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
)

type StructInterfaceSource struct {
	node *remote.Node
	impl api.StructInterface
}

var _ remote.IObjectSource = (*StructInterfaceSource)(nil)
var _ api.INotifier = (*StructInterfaceSource)(nil)

func NewStructInterfaceSource() *StructInterfaceSource {
	return &StructInterfaceSource{}
}

func (s *StructInterfaceSource) SetImplementation(impl api.StructInterface) {
	s.impl = impl
}

func (s *StructInterfaceSource) ObjectId() string {
	return "testbed1.StructInterface"
}

func (s *StructInterfaceSource) Invoke(methodId string, args core.Args) (core.Any, error) {
	log.Printf("source: invoke: %s %v", methodId, args)
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	name := core.SymbolIdToMember(methodId)
	switch name {
	case "funcBool":
		paramBool, err := api.AsStructBool(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.FuncBool(paramBool), nil
	case "funcInt":
		paramInt, err := api.AsStructInt(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.FuncInt(paramInt), nil
	case "funcFloat":
		paramFloat, err := api.AsStructFloat(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.FuncFloat(paramFloat), nil
	case "funcString":
		paramString, err := api.AsStructString(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.FuncString(paramString), nil
	default:
		return nil, fmt.Errorf("unknown method: %s", name)
	}
}

func (s *StructInterfaceSource) SetProperty(propertyId string, value core.Any) error {
	if s.impl == nil {
		return fmt.Errorf("no implementation")
	}
	name := core.SymbolIdToMember(propertyId)
	switch name {
	case "propBool":
		prop, err := api.AsStructBool(value)
		if err != nil {
			return err
		}
		s.impl.SetPropBool(prop)
	case "propInt":
		prop, err := api.AsStructInt(value)
		if err != nil {
			return err
		}
		s.impl.SetPropInt(prop)
	case "propFloat":
		prop, err := api.AsStructFloat(value)
		if err != nil {
			return err
		}
		s.impl.SetPropFloat(prop)
	case "propString":
		prop, err := api.AsStructString(value)
		if err != nil {
			return err
		}
		s.impl.SetPropString(prop)
	default:
		return fmt.Errorf("StructInterfaceSource.SetProperty: unknown property %s", propertyId)
	}
	return nil
}

func (s *StructInterfaceSource) CollectProperties() (core.KWArgs, error) {
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

func (s *StructInterfaceSource) Linked(objectId string, node *remote.Node) error {
	if objectId != s.ObjectId() {
		return fmt.Errorf("not my object: %s", objectId)
	}
	s.node = node
	return nil
}

func (s *StructInterfaceSource) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *StructInterfaceSource) NotifySignal(name string, args []any) {
	signalId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}
