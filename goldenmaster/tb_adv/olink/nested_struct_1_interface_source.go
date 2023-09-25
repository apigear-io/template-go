package olink

import (
	"fmt"
	"goldenmaster/tb_adv/api"
	"log"
	"olink/pkg/core"
	"olink/pkg/remote"
)

type NestedStruct1InterfaceSource struct {
	node *remote.Node
	impl api.NestedStruct1Interface
}

var _ remote.IObjectSource = (*NestedStruct1InterfaceSource)(nil)
var _ api.INotifier = (*NestedStruct1InterfaceSource)(nil)

func NewNestedStruct1InterfaceSource() *NestedStruct1InterfaceSource {
	return &NestedStruct1InterfaceSource{}
}

func (s *NestedStruct1InterfaceSource) SetImplementation(impl api.NestedStruct1Interface) {
	s.impl = impl
}

func (s *NestedStruct1InterfaceSource) ObjectId() string {
	return "tb.adv.NestedStruct1Interface"
}

func (s *NestedStruct1InterfaceSource) Invoke(methodId string, args core.Args) (core.Any, error) {
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

func (s *NestedStruct1InterfaceSource) SetProperty(propertyId string, value core.Any) error {
	if s.impl == nil {
		return fmt.Errorf("no implementation")
	}
	name := core.ToMember(propertyId)
	switch name {
	case "prop1":
		prop, err := api.AsNestedStruct1(value)
		if err != nil {
			return err
		}
		s.impl.SetProp1(prop)
	default:
		return fmt.Errorf("NestedStruct1InterfaceSource.SetProperty: unknown property %s", propertyId)
	}
	return nil
}

func (s *NestedStruct1InterfaceSource) CollectProperties() (core.KWArgs, error) {
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	return core.KWArgs{
		"prop1": s.impl.GetProp1(),
	}, nil
}

func (s *NestedStruct1InterfaceSource) Linked(objectId string, node *remote.Node) error {
	if objectId != s.ObjectId() {
		return fmt.Errorf("not my object: %s", objectId)
	}
	s.node = node
	return nil
}

func (s *NestedStruct1InterfaceSource) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *NestedStruct1InterfaceSource) NotifySignal(name string, args []any) {
	signalId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}
