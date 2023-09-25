package olink

import (
	"fmt"
	"goldenmaster/tb_conflict/api"
	"log"
	"olink/pkg/core"
	"olink/pkg/remote"
)

type Conflict4Source struct {
	node *remote.Node
	impl api.Conflict4
}

var _ remote.IObjectSource = (*Conflict4Source)(nil)
var _ api.INotifier = (*Conflict4Source)(nil)

func NewConflict4Source() *Conflict4Source {
	return &Conflict4Source{}
}

func (s *Conflict4Source) SetImplementation(impl api.Conflict4) {
	s.impl = impl
}

func (s *Conflict4Source) ObjectId() string {
	return "tb.conflict.Conflict4"
}

func (s *Conflict4Source) Invoke(methodId string, args core.Args) (core.Any, error) {
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

func (s *Conflict4Source) SetProperty(propertyId string, value core.Any) error {
	if s.impl == nil {
		return fmt.Errorf("no implementation")
	}
	name := core.ToMember(propertyId)
	switch name {
	default:
		return fmt.Errorf("Conflict4Source.SetProperty: unknown property %s", propertyId)
	}
	return nil
}

func (s *Conflict4Source) CollectProperties() (core.KWArgs, error) {
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	return core.KWArgs{}, nil
}

func (s *Conflict4Source) Linked(objectId string, node *remote.Node) error {
	if objectId != s.ObjectId() {
		return fmt.Errorf("not my object: %s", objectId)
	}
	s.node = node
	return nil
}

func (s *Conflict4Source) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *Conflict4Source) NotifySignal(name string, args []any) {
	signalId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}
