package olink

import (
	"fmt"
	"log"
	"testbed/tb_same2/api"

	"github.com/apigear-io/objectlink-core-go/olink/core"
	"github.com/apigear-io/objectlink-core-go/olink/remote"
)

type SameEnum1InterfaceSource struct {
	node *remote.Node
	impl api.SameEnum1Interface
}

var _ remote.IObjectSource = (*SameEnum1InterfaceSource)(nil)
var _ api.INotifier = (*SameEnum1InterfaceSource)(nil)

func NewSameEnum1InterfaceSource() *SameEnum1InterfaceSource {
	return &SameEnum1InterfaceSource{}
}

func (s *SameEnum1InterfaceSource) SetImplementation(impl api.SameEnum1Interface) {
	s.impl = impl
}

func (s *SameEnum1InterfaceSource) ObjectId() string {
	return "tb.same2.SameEnum1Interface"
}

func (s *SameEnum1InterfaceSource) Invoke(methodId string, args core.Args) (core.Any, error) {
	log.Printf("source: invoke: %s %v", methodId, args)
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	name := core.SymbolIdToMember(methodId)
	switch name {
	case "func1":
		param1, err := api.AsEnum1(args[0])
		if err != nil {
			return nil, err
		}
		return s.impl.Func1(param1), nil
	default:
		return nil, fmt.Errorf("unknown method: %s", name)
	}
}

func (s *SameEnum1InterfaceSource) SetProperty(propertyId string, value core.Any) error {
	if s.impl == nil {
		return fmt.Errorf("no implementation")
	}
	name := core.SymbolIdToMember(propertyId)
	switch name {
	case "prop1":
		prop, err := api.AsEnum1(value)
		if err != nil {
			return err
		}
		s.impl.SetProp1(prop)
	default:
		return fmt.Errorf("SameEnum1InterfaceSource.SetProperty: unknown property %s", propertyId)
	}
	return nil
}

func (s *SameEnum1InterfaceSource) CollectProperties() (core.KWArgs, error) {
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	return core.KWArgs{
		"prop1": s.impl.GetProp1(),
	}, nil
}

func (s *SameEnum1InterfaceSource) Linked(objectId string, node *remote.Node) error {
	if objectId != s.ObjectId() {
		return fmt.Errorf("not my object: %s", objectId)
	}
	s.node = node
	return nil
}

func (s *SameEnum1InterfaceSource) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *SameEnum1InterfaceSource) NotifySignal(name string, args []any) {
	signalId := core.MakeSymbolId(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}
