package olink

import (
	"fmt"
	"log"
	"testbed/testbed2/api"

	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
)

type NestedStruct3InterfaceSource struct {
	node *remote.Node
	impl api.NestedStruct3Interface
}

var _ remote.IObjectSource = (*NestedStruct3InterfaceSource)(nil)
var _ api.INotifier = (*NestedStruct3InterfaceSource)(nil)

func NewNestedStruct3InterfaceSource() *NestedStruct3InterfaceSource {
	return &NestedStruct3InterfaceSource{}
}

func (s *NestedStruct3InterfaceSource) SetImplementation(impl api.NestedStruct3Interface) {
	s.impl = impl
}

func (s *NestedStruct3InterfaceSource) ObjectId() string {
	return "testbed2.NestedStruct3Interface"
}

func (s *NestedStruct3InterfaceSource) Invoke(methodId string, args core.Args) (core.Any, error) {
	log.Printf("source: invoke: %s %v", methodId, args)
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	name := core.SymbolIdToMember(methodId)
	switch name {
	case "func1":
		param1, err := api.AsNestedStruct1(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.Func1(param1), nil
	case "func2":
		param1, err := api.AsNestedStruct1(args[0])
		if err != nil {
			return nil, err
		}
		param2, err := api.AsNestedStruct2(args[1])
		if err != nil {
			return nil, err
		}
		return s.impl.Func2(param1, param2), nil
	case "func3":
		param1, err := api.AsNestedStruct1(args[0])
		if err != nil {
			return nil, err
		}
		param2, err := api.AsNestedStruct2(args[1])
		if err != nil {
			return nil, err
		}
		param3, err := api.AsNestedStruct3(args[2])
		if err != nil {
			return nil, err
		}
		return s.impl.Func3(param1, param2, param3), nil
	default:
		return nil, fmt.Errorf("unknown method: %s", name)
	}
}

func (s *NestedStruct3InterfaceSource) SetProperty(propertyId string, value core.Any) error {
	if s.impl == nil {
		return fmt.Errorf("no implementation")
	}
	name := core.SymbolIdToMember(propertyId)
	switch name {
	case "prop1":
		prop, err := api.AsNestedStruct1(value)
		if err != nil {
			return err
		}
		s.impl.SetProp1(prop)
	case "prop2":
		prop, err := api.AsNestedStruct2(value)
		if err != nil {
			return err
		}
		s.impl.SetProp2(prop)
	case "prop3":
		prop, err := api.AsNestedStruct3(value)
		if err != nil {
			return err
		}
		s.impl.SetProp3(prop)
	default:
		return fmt.Errorf("NestedStruct3InterfaceSource.SetProperty: unknown property %s", propertyId)
	}
	return nil
}

func (s *NestedStruct3InterfaceSource) CollectProperties() (core.KWArgs, error) {
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	return core.KWArgs{
		"prop1": s.impl.GetProp1(),
		"prop2": s.impl.GetProp2(),
		"prop3": s.impl.GetProp3(),
	}, nil
}

func (s *NestedStruct3InterfaceSource) Linked(objectId string, node *remote.Node) error {
	if objectId != s.ObjectId() {
		return fmt.Errorf("not my object: %s", objectId)
	}
	s.node = node
	return nil
}

func (s *NestedStruct3InterfaceSource) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *NestedStruct3InterfaceSource) NotifySignal(name string, args []any) {
	signalId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}
