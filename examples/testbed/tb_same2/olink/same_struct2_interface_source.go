package olink

import (
	"fmt"
	"log"
	"testbed/tb_same2/api"

	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
)

type SameStruct2InterfaceSource struct {
	node *remote.Node
	impl api.SameStruct2Interface
}

var _ remote.IObjectSource = (*SameStruct2InterfaceSource)(nil)
var _ api.INotifier = (*SameStruct2InterfaceSource)(nil)

func NewSameStruct2InterfaceSource() *SameStruct2InterfaceSource {
	return &SameStruct2InterfaceSource{}
}

func (s *SameStruct2InterfaceSource) SetImplementation(impl api.SameStruct2Interface) {
	s.impl = impl
}

func (s *SameStruct2InterfaceSource) ObjectId() string {
	return "tb.same2.SameStruct2Interface"
}

func (s *SameStruct2InterfaceSource) Invoke(methodId string, args core.Args) (core.Any, error) {
	log.Printf("source: invoke: %s %v", methodId, args)
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	name := core.SymbolIdToMember(methodId)
	switch name {
	case "func1":
		param1, err := api.AsStruct1(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.Func1(param1), nil
	case "func2":
		param1, err := api.AsStruct1(args[0])
		if err != nil {
			return nil, err
		}
		param2, err := api.AsStruct2(args[1])
		if err != nil {
			return nil, err
		}
		return s.impl.Func2(param1, param2), nil
	default:
		return nil, fmt.Errorf("unknown method: %s", name)
	}
}

func (s *SameStruct2InterfaceSource) SetProperty(propertyId string, value core.Any) error {
	if s.impl == nil {
		return fmt.Errorf("no implementation")
	}
	name := core.SymbolIdToMember(propertyId)
	switch name {
	case "prop1":
		prop, err := api.AsStruct1(value)
		if err != nil {
			return err
		}
		s.impl.SetProp1(prop)
	case "prop2":
		prop, err := api.AsStruct2(value)
		if err != nil {
			return err
		}
		s.impl.SetProp2(prop)
	default:
		return fmt.Errorf("SameStruct2InterfaceSource.SetProperty: unknown property %s", propertyId)
	}
	return nil
}

func (s *SameStruct2InterfaceSource) CollectProperties() (core.KWArgs, error) {
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	return core.KWArgs{
		"prop1": s.impl.GetProp1(),
		"prop2": s.impl.GetProp2(),
	}, nil
}

func (s *SameStruct2InterfaceSource) Linked(objectId string, node *remote.Node) error {
	if objectId != s.ObjectId() {
		return fmt.Errorf("not my object: %s", objectId)
	}
	s.node = node
	return nil
}

func (s *SameStruct2InterfaceSource) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *SameStruct2InterfaceSource) NotifySignal(name string, args []any) {
	signalId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}
