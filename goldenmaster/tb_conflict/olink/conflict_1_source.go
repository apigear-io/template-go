package olink

import (
	"fmt"
	"goldenmaster/tb_conflict/api"
	"log"
	"olink/pkg/core"
	"olink/pkg/remote"
)

type Conflict1Source struct {
	node *remote.Node
	impl api.Conflict1
}

var _ remote.IObjectSource = (*Conflict1Source)(nil)
var _ api.INotifier = (*Conflict1Source)(nil)

func NewConflict1Source() *Conflict1Source {
	return &Conflict1Source{}
}

func (s *Conflict1Source) SetImplementation(impl api.Conflict1) {
	s.impl = impl
}

func (s *Conflict1Source) ObjectId() string {
	return "tb.conflict.Conflict1"
}

func (s *Conflict1Source) Invoke(methodId string, args core.Args) (core.Any, error) {
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

func (s *Conflict1Source) SetProperty(propertyId string, value core.Any) error {
	if s.impl == nil {
		return fmt.Errorf("no implementation")
	}
	name := core.ToMember(propertyId)
	switch name {
	case "sameName":
		prop, err := api.AsInt(value)
		if err != nil {
			return err
		}
		s.impl.SetSameName(prop)
	default:
		return fmt.Errorf("Conflict1Source.SetProperty: unknown property %s", propertyId)
	}
	return nil
}

func (s *Conflict1Source) CollectProperties() (core.KWArgs, error) {
	if s.impl == nil {
		return nil, fmt.Errorf("no implementation")
	}
	return core.KWArgs{
		"sameName": s.impl.GetSameName(),
	}, nil
}

func (s *Conflict1Source) Linked(objectId string, node *remote.Node) error {
	if objectId != s.ObjectId() {
		return fmt.Errorf("not my object: %s", objectId)
	}
	s.node = node
	return nil
}

func (s *Conflict1Source) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *Conflict1Source) NotifySignal(name string, args []any) {
	signalId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}
