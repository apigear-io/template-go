package olink

import (
	"fmt"
	"log"
	"testbed/testbed1/api"

	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
)

type StructArrayInterfaceSource struct {
	node *remote.Node
	impl api.StructArrayInterface
}

var _ remote.IObjectSource = (*StructArrayInterfaceSource)(nil)
var _ api.INotifier = (*StructArrayInterfaceSource)(nil)

func NewStructArrayInterfaceSource() *StructArrayInterfaceSource {
	return &StructArrayInterfaceSource{}
}

func (s *StructArrayInterfaceSource) SetImplementation(impl api.StructArrayInterface) {
	s.impl = impl
}

func (s *StructArrayInterfaceSource) ObjectId() string {
	return "testbed1.StructArrayInterface"
}

func (s *StructArrayInterfaceSource) Invoke(methodId string, args core.Args) (core.Any, error) {
	log.Printf("source: invoke: %s %v", methodId, args)
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	name := core.SymbolIdToMember(methodId)
	switch name {
	case "funcBool":
		paramBool, err := api.AsStructBoolArray(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.FuncBool(paramBool), nil
	case "funcInt":
		paramInt, err := api.AsStructIntArray(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.FuncInt(paramInt), nil
	case "funcFloat":
		paramFloat, err := api.AsStructFloatArray(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.FuncFloat(paramFloat), nil
	case "funcString":
		paramString, err := api.AsStructStringArray(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.FuncString(paramString), nil
	default:
		return nil, fmt.Errorf("unknown method: %s", name)
	}
}

func (s *StructArrayInterfaceSource) SetProperty(propertyId string, value core.Any) error {
	if s.impl == nil {
		return fmt.Errorf("no implementation")
	}
	name := core.SymbolIdToMember(propertyId)
	switch name {
	case "propBool":
		prop, err := api.AsStructBoolArray(value)
		if err != nil {
			return err
		}
		s.impl.SetPropBool(prop)
	case "propInt":
		prop, err := api.AsStructIntArray(value)
		if err != nil {
			return err
		}
		s.impl.SetPropInt(prop)
	case "propFloat":
		prop, err := api.AsStructFloatArray(value)
		if err != nil {
			return err
		}
		s.impl.SetPropFloat(prop)
	case "propString":
		prop, err := api.AsStructStringArray(value)
		if err != nil {
			return err
		}
		s.impl.SetPropString(prop)
	default:
		return fmt.Errorf("StructArrayInterfaceSource.SetProperty: unknown property %s", propertyId)
	}
	return nil
}

func (s *StructArrayInterfaceSource) CollectProperties() (core.KWArgs, error) {
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

func (s *StructArrayInterfaceSource) Linked(objectId string, node *remote.Node) error {
	if objectId != s.ObjectId() {
		return fmt.Errorf("not my object: %s", objectId)
	}
	s.node = node
	return nil
}

func (s *StructArrayInterfaceSource) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *StructArrayInterfaceSource) NotifySignal(name string, args []any) {
	signalId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}
