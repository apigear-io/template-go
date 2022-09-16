package olink

import (
    "fmt"
    "log"
	"olink/pkg/core"
    "olink/pkg/remote"
    "goldenmaster/tb_conflict/api"
)

type Conflict2Source struct {
    node *remote.Node
    impl api.Conflict2
}

var _ remote.IObjectSource = (*Conflict2Source)(nil)
var _ api.INotifier = (*Conflict2Source)(nil)

func NewConflict2Source() *Conflict2Source {
    return &Conflict2Source{}
}

func (s *Conflict2Source) SetImplementation(impl api.Conflict2) {
    s.impl = impl
}

func (s *Conflict2Source) ObjectId() string {
	return "tb.conflict.Conflict2"
}

func (s *Conflict2Source) Invoke(methodId string, args core.Args) (core.Any, error) {
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

func (s *Conflict2Source) SetProperty(propertyId string, value core.Any) error {
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
        return fmt.Errorf("Conflict2Source.SetProperty: unknown property %s", propertyId)
    }
    return nil
}

func (s *Conflict2Source) CollectProperties() (core.KWArgs, error) {
    if s.impl == nil {
        return nil, fmt.Errorf("no implementation")
    }
    return core.KWArgs{
        "sameName": s.impl.GetSameName(),
    }, nil
}


func (s *Conflict2Source) Linked(objectId string, node *remote.Node) error {
    if objectId != s.ObjectId() {
        return fmt.Errorf("not my object: %s", objectId)
    }
    s.node = node
    return nil
}

func (s *Conflict2Source) NotifyPropertyChanged(name string, value any) {
	propertyId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifyPropertyChange(propertyId, value)
}

func (s *Conflict2Source) NotifySignal(name string, args []any) {
	signalId := core.MakeIdentifier(s.ObjectId(), name)
	s.node.NotifySignal(signalId, args)
}
