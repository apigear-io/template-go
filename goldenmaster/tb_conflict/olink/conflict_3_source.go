package olink

import (
    "fmt"
    "log"
	"olink/pkg/core"
    "olink/pkg/remote"
    "goldenmaster/tb_conflict/api"
)

type Conflict3Source struct {
    node *remote.Node
    impl api.Conflict3
}

var _ remote.IObjectSource = (*Conflict3Source)(nil)
var _ api.INotifier = (*Conflict3Source)(nil)

func NewConflict3Source() *Conflict3Source {
    return &Conflict3Source{}
}

func (s *Conflict3Source) SetImplementation(impl api.Conflict3) {
    s.impl = impl
}

func (s *Conflict3Source) ObjectId() string {
	return "tb.conflict.Conflict3"
}

func (s *Conflict3Source) Invoke(methodId string, args core.Args) (core.Any, error) {
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

func (s *Conflict3Source) SetProperty(propertyId string, value core.Any) error {
    if s.impl == nil {
        return fmt.Errorf("no implementation")
    }
    name := core.ToMember(propertyId)
    switch name {
    default:
        return fmt.Errorf("Conflict3Source.SetProperty: unknown property %s", propertyId)
    }
    return nil
}

func (s *Conflict3Source) CollectProperties() (core.KWArgs, error) {
    if s.impl == nil {
        return nil, fmt.Errorf("no implementation")
    }
    return core.KWArgs{
    }, nil
}


func (s *Conflict3Source) Linked(objectId string, node *remote.Node) error {
    if objectId != s.ObjectId() {
        return fmt.Errorf("not my object: %s", objectId)
    }
    s.node = node
    return nil
}

func (s *Conflict3Source) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *Conflict3Source) NotifySignal(name string, args []any) {
	signalId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}
