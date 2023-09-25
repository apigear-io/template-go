package olink

import (
	"fmt"
	"goldenmaster/tb_simple/api"
	"log"
	"olink/pkg/core"
	"olink/pkg/remote"
)

type SimpleArrayInterfaceSource struct {
	node *remote.Node
	impl api.SimpleArrayInterface
}

var _ remote.IObjectSource = (*SimpleArrayInterfaceSource)(nil)
var _ api.INotifier = (*SimpleArrayInterfaceSource)(nil)

func NewSimpleArrayInterfaceSource() *SimpleArrayInterfaceSource {
	return &SimpleArrayInterfaceSource{}
}

func (s *SimpleArrayInterfaceSource) SetImplementation(impl api.SimpleArrayInterface) {
	s.impl = impl
}

func (s *SimpleArrayInterfaceSource) ObjectId() string {
	return "tb.simple.SimpleArrayInterface"
}

func (s *SimpleArrayInterfaceSource) Invoke(methodId string, args core.Args) (core.Any, error) {
	log.Printf("source: invoke: %s %v", methodId, args)
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	name := core.ToMember(methodId)
	switch name {
	case "funcBool":
		paramBool, err := api.AsBoolArray(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.FuncBool(paramBool), nil
	case "funcInt":
		paramInt, err := api.AsIntArray(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.FuncInt(paramInt), nil
	case "funcFloat":
		paramFloat, err := api.AsFloatArray(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.FuncFloat(paramFloat), nil
	case "funcString":
		paramString, err := api.AsStringArray(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.FuncString(paramString), nil
	default:
		return nil, fmt.Errorf("unknown method: %s", name)
	}
}

func (s *SimpleArrayInterfaceSource) SetProperty(propertyId string, value core.Any) error {
	if s.impl == nil {
		return fmt.Errorf("no implementation")
	}
	name := core.ToMember(propertyId)
	switch name {
	case "propBool":
		prop, err := api.AsBoolArray(value)
		if err != nil {
			return err
		}
		s.impl.SetPropBool(prop)
	case "propInt":
		prop, err := api.AsIntArray(value)
		if err != nil {
			return err
		}
		s.impl.SetPropInt(prop)
	case "propFloat":
		prop, err := api.AsFloatArray(value)
		if err != nil {
			return err
		}
		s.impl.SetPropFloat(prop)
	case "propString":
		prop, err := api.AsStringArray(value)
		if err != nil {
			return err
		}
		s.impl.SetPropString(prop)
	default:
		return fmt.Errorf("SimpleArrayInterfaceSource.SetProperty: unknown property %s", propertyId)
	}
	return nil
}

func (s *SimpleArrayInterfaceSource) CollectProperties() (core.KWArgs, error) {
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

func (s *SimpleArrayInterfaceSource) Linked(objectId string, node *remote.Node) error {
	if objectId != s.ObjectId() {
		return fmt.Errorf("not my object: %s", objectId)
	}
	s.node = node
	return nil
}

func (s *SimpleArrayInterfaceSource) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *SimpleArrayInterfaceSource) NotifySignal(name string, args []any) {
	signalId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}
