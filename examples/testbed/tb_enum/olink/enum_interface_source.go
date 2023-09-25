package olink

import (
	"fmt"
	"log"
	"testbed/tb_enum/api"

	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
)

type EnumInterfaceSource struct {
	node *remote.Node
	impl api.EnumInterface
}

var _ remote.IObjectSource = (*EnumInterfaceSource)(nil)
var _ api.INotifier = (*EnumInterfaceSource)(nil)

func NewEnumInterfaceSource() *EnumInterfaceSource {
	return &EnumInterfaceSource{}
}

func (s *EnumInterfaceSource) SetImplementation(impl api.EnumInterface) {
	s.impl = impl
}

func (s *EnumInterfaceSource) ObjectId() string {
	return "tb.enum.EnumInterface"
}

func (s *EnumInterfaceSource) Invoke(methodId string, args core.Args) (core.Any, error) {
	log.Printf("source: invoke: %s %v", methodId, args)
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	name := core.SymbolIdToMember(methodId)
	switch name {
	case "func0":
		param0, err := api.AsEnum0(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.Func0(param0), nil
	case "func1":
		param1, err := api.AsEnum1(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.Func1(param1), nil
	case "func2":
		param2, err := api.AsEnum2(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.Func2(param2), nil
	case "func3":
		param3, err := api.AsEnum3(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.Func3(param3), nil
	default:
		return nil, fmt.Errorf("unknown method: %s", name)
	}
}

func (s *EnumInterfaceSource) SetProperty(propertyId string, value core.Any) error {
	if s.impl == nil {
		return fmt.Errorf("no implementation")
	}
	name := core.SymbolIdToMember(propertyId)
	switch name {
	case "prop0":
		prop, err := api.AsEnum0(value)
		if err != nil {
			return err
		}
		s.impl.SetProp0(prop)
	case "prop1":
		prop, err := api.AsEnum1(value)
		if err != nil {
			return err
		}
		s.impl.SetProp1(prop)
	case "prop2":
		prop, err := api.AsEnum2(value)
		if err != nil {
			return err
		}
		s.impl.SetProp2(prop)
	case "prop3":
		prop, err := api.AsEnum3(value)
		if err != nil {
			return err
		}
		s.impl.SetProp3(prop)
	default:
		return fmt.Errorf("EnumInterfaceSource.SetProperty: unknown property %s", propertyId)
	}
	return nil
}

func (s *EnumInterfaceSource) CollectProperties() (core.KWArgs, error) {
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	return core.KWArgs{
		"prop0": s.impl.GetProp0(),
		"prop1": s.impl.GetProp1(),
		"prop2": s.impl.GetProp2(),
		"prop3": s.impl.GetProp3(),
	}, nil
}

func (s *EnumInterfaceSource) Linked(objectId string, node *remote.Node) error {
	if objectId != s.ObjectId() {
		return fmt.Errorf("not my object: %s", objectId)
	}
	s.node = node
	return nil
}

func (s *EnumInterfaceSource) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *EnumInterfaceSource) NotifySignal(name string, args []any) {
	signalId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}
