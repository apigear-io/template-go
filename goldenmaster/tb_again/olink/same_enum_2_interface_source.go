package olink

import (
	"fmt"
	"goldenmaster/tb_again/api"
	"log"
	"olink/pkg/core"
	"olink/pkg/remote"
)

type SameEnum2InterfaceSource struct {
	node *remote.Node
	impl api.SameEnum2Interface
}

var _ remote.IObjectSource = (*SameEnum2InterfaceSource)(nil)
var _ api.INotifier = (*SameEnum2InterfaceSource)(nil)

func NewSameEnum2InterfaceSource() *SameEnum2InterfaceSource {
	return &SameEnum2InterfaceSource{}
}

func (s *SameEnum2InterfaceSource) SetImplementation(impl api.SameEnum2Interface) {
	s.impl = impl
}

func (s *SameEnum2InterfaceSource) ObjectId() string {
	return "tb.again.SameEnum2Interface"
}

func (s *SameEnum2InterfaceSource) Invoke(methodId string, args core.Args) (core.Any, error) {
	log.Printf("source: invoke: %s %v", methodId, args)
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	name := core.ToMember(methodId)
	switch name {
	default:
		return nil, fmt.Errorf("unknown method: %s", name)
	}
}

func (s *SameEnum2InterfaceSource) SetProperty(propertyId string, value core.Any) error {
	if s.impl == nil {
		return fmt.Errorf("no implementation")
	}
	name := core.ToMember(propertyId)
	switch name {
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
	default:
		return fmt.Errorf("SameEnum2InterfaceSource.SetProperty: unknown property %s", propertyId)
	}
	return nil
}

func (s *SameEnum2InterfaceSource) CollectProperties() (core.KWArgs, error) {
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	return core.KWArgs{
		"prop1": s.impl.GetProp1(),
		"prop2": s.impl.GetProp2(),
	}, nil
}

func (s *SameEnum2InterfaceSource) Linked(objectId string, node *remote.Node) error {
	if objectId != s.ObjectId() {
		return fmt.Errorf("not my object: %s", objectId)
	}
	s.node = node
	return nil
}

func (s *SameEnum2InterfaceSource) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *SameEnum2InterfaceSource) NotifySignal(name string, args []any) {
	signalId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}
