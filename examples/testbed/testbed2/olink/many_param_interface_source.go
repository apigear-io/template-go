package olink

import (
	"fmt"
	"log"
	"testbed/testbed2/api"

	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
)

type ManyParamInterfaceSource struct {
	node *remote.Node
	impl api.ManyParamInterface
}

var _ remote.IObjectSource = (*ManyParamInterfaceSource)(nil)
var _ api.INotifier = (*ManyParamInterfaceSource)(nil)

func NewManyParamInterfaceSource() *ManyParamInterfaceSource {
	return &ManyParamInterfaceSource{}
}

func (s *ManyParamInterfaceSource) SetImplementation(impl api.ManyParamInterface) {
	s.impl = impl
}

func (s *ManyParamInterfaceSource) ObjectId() string {
	return "testbed2.ManyParamInterface"
}

func (s *ManyParamInterfaceSource) Invoke(methodId string, args core.Args) (core.Any, error) {
	log.Printf("source: invoke: %s %v", methodId, args)
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	name := core.SymbolIdToMember(methodId)
	switch name {
	case "func1":
		param1, err := api.AsInt(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.Func1(param1), nil
	case "func2":
		param1, err := api.AsInt(args[0])
		if err != nil {
			return nil, err
		}
		param2, err := api.AsInt(args[1])
		if err != nil {
			return nil, err
		}
		return s.impl.Func2(param1, param2), nil
	case "func3":
		param1, err := api.AsInt(args[0])
		if err != nil {
			return nil, err
		}
		param2, err := api.AsInt(args[1])
		if err != nil {
			return nil, err
		}
		param3, err := api.AsInt(args[2])
		if err != nil {
			return nil, err
		}
		return s.impl.Func3(param1, param2, param3), nil
	case "func4":
		param1, err := api.AsInt(args[0])
		if err != nil {
			return nil, err
		}
		param2, err := api.AsInt(args[1])
		if err != nil {
			return nil, err
		}
		param3, err := api.AsInt(args[2])
		if err != nil {
			return nil, err
		}
		param4, err := api.AsInt(args[3])
		if err != nil {
			return nil, err
		}
		return s.impl.Func4(param1, param2, param3, param4), nil
	default:
		return nil, fmt.Errorf("unknown method: %s", name)
	}
}

func (s *ManyParamInterfaceSource) SetProperty(propertyId string, value core.Any) error {
	if s.impl == nil {
		return fmt.Errorf("no implementation")
	}
	name := core.SymbolIdToMember(propertyId)
	switch name {
	case "prop1":
		prop, err := api.AsInt(value)
		if err != nil {
			return err
		}
		s.impl.SetProp1(prop)
	case "prop2":
		prop, err := api.AsInt(value)
		if err != nil {
			return err
		}
		s.impl.SetProp2(prop)
	case "prop3":
		prop, err := api.AsInt(value)
		if err != nil {
			return err
		}
		s.impl.SetProp3(prop)
	case "prop4":
		prop, err := api.AsInt(value)
		if err != nil {
			return err
		}
		s.impl.SetProp4(prop)
	default:
		return fmt.Errorf("ManyParamInterfaceSource.SetProperty: unknown property %s", propertyId)
	}
	return nil
}

func (s *ManyParamInterfaceSource) CollectProperties() (core.KWArgs, error) {
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	return core.KWArgs{
		"prop1": s.impl.GetProp1(),
		"prop2": s.impl.GetProp2(),
		"prop3": s.impl.GetProp3(),
		"prop4": s.impl.GetProp4(),
	}, nil
}

func (s *ManyParamInterfaceSource) Linked(objectId string, node *remote.Node) error {
	if objectId != s.ObjectId() {
		return fmt.Errorf("not my object: %s", objectId)
	}
	s.node = node
	return nil
}

func (s *ManyParamInterfaceSource) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *ManyParamInterfaceSource) NotifySignal(name string, args []any) {
	signalId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}
